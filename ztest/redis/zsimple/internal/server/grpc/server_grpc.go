package grpc

import (
	"context"
	"fmt"
	"math"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	pb "class/ztest/redis/zsimple/api"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/go-kratos/kratos/pkg/log"
	nmd "github.com/go-kratos/kratos/pkg/net/metadata"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden/ratelimiter"
	"github.com/go-kratos/kratos/pkg/net/trace"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	_abortIndex int8 = math.MaxInt8 / 2
)

type Server struct {
	conf     *ServerConfig
	mutex    sync.RWMutex
	server   *grpc.Server
	handlers []grpc.UnaryServerInterceptor
}

type ServerConfig struct {
	Network string         `dsn:"network"`
	Addr    string         `dsn:"address"`
	Timeout xtime.Duration `dsn:"query.timeout"`
}

func New2(svc pb.DemoServer) (gs *Server, err error) {
	var (
		cfg ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	gs = NewServer(&cfg)
	pb.RegisterDemoServer(gs.server, svc)
	gs, err = gs.Start()
	return
}

func NewServer(conf *ServerConfig, opt ...grpc.ServerOption) (s *Server) {
	s = new(Server)
	s.SetConfig(conf)
	s.server = grpc.NewServer(opt...)
	s.Use(s.recovery(), s.handle(), )
	s.Use(ratelimiter.New(nil).Limit())
	return
}

// SetConfig hot reloads server config
func (s *Server) SetConfig(conf *ServerConfig) {
	if conf == nil {
		conf = &ServerConfig{}
	}
	if conf.Timeout <= 0 {
		conf.Timeout = xtime.Duration(time.Second)
	}
	if conf.Addr == "" {
		conf.Addr = "0.0.0.0:9000"
	}
	if conf.Network == "" {
		conf.Network = "tcp"
	}
	s.mutex.Lock()
	s.conf = conf
	s.mutex.Unlock()
	return
}

func (s *Server) Start() (*Server, error) {
	_, err := s.startWithAddr()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Server) StartWithAddr() (*Server, net.Addr, error) {
	addr, err := s.startWithAddr()
	if err != nil {
		return nil, nil, err
	}
	return s, addr, nil
}

func (s *Server) startWithAddr() (net.Addr, error) {
	lis, err := net.Listen(s.conf.Network, s.conf.Addr)
	if err != nil {
		return nil, err
	}
	log.Info("warden2: start grpc listen addr: %v", lis.Addr())
	reflection.Register(s.server)
	go func() {
		if err := s.server.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return lis.Addr(), nil
}

func (s *Server) Use(handlers ...grpc.UnaryServerInterceptor) *Server {
	finalSize := len(s.handlers) + len(handlers)
	if finalSize >= int(_abortIndex) {
		panic("warden: server use too many handlers")
	}
	mergedHandlers := make([]grpc.UnaryServerInterceptor, finalSize)
	copy(mergedHandlers, s.handlers)
	copy(mergedHandlers[len(s.handlers):], handlers)
	s.handlers = mergedHandlers
	return s
}

// handle return a new unary server interceptor for OpenTracing\Logging\LinkTimeout.
func (s *Server) handle() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var (
			cancel func()
			addr   string
		)
		s.mutex.RLock()
		conf := s.conf
		s.mutex.RUnlock()
		// get derived timeout from grpc context,
		// compare with the warden configured,
		// and use the minimum one
		timeout := time.Duration(conf.Timeout)
		if dl, ok := ctx.Deadline(); ok {
			ctimeout := time.Until(dl)
			if ctimeout-time.Millisecond*20 > 0 {
				ctimeout = ctimeout - time.Millisecond*20
			}
			if timeout > ctimeout {
				timeout = ctimeout
			}
		}
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()

		// get grpc metadata(trace & remote_ip & color)
		var t trace.Trace
		cmd := nmd.MD{}
		if gmd, ok := metadata.FromIncomingContext(ctx); ok {
			t, _ = trace.Extract(trace.GRPCFormat, gmd)
			for key, vals := range gmd {
				if nmd.IsIncomingKey(key) {
					cmd[key] = vals[0]
				}
			}
		}
		if t == nil {
			t = trace.New(args.FullMethod)
		} else {
			t.SetTitle(args.FullMethod)
		}

		if pr, ok := peer.FromContext(ctx); ok {
			addr = pr.Addr.String()
			t.SetTag(trace.String(trace.TagAddress, addr))
		}
		defer t.Finish(&err)

		// use common meta data context instead of grpc context
		ctx = nmd.NewContext(ctx, cmd)
		ctx = trace.NewContext(ctx, t)

		resp, err = handler(ctx, req)
		return resp, err
	}
}

// recovery is a server interceptor that recovers from any panics.
func (s *Server) recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if rerr := recover(); rerr != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				rs := runtime.Stack(buf, false)
				if rs > size {
					rs = size
				}
				buf = buf[:rs]
				pl := fmt.Sprintf("grpc server panic: %v\n%v\n%s\n", req, rerr, buf)
				fmt.Fprint(os.Stderr, pl)
				log.Error(pl)
				err = status.Errorf(codes.Unknown, ecode.ServerErr.Error())
			}
		}()
		resp, err = handler(ctx, req)
		return
	}
}
