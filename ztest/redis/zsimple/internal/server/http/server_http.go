package http

import (
	"sync"
	"time"

	pb "class/ztest/redis/zsimple/api"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"google.golang.org/grpc"
)

type Server struct {
	conf     *ServerConfig
	mutex    sync.RWMutex
	engine   *gin.Engine
	handlers []grpc.UnaryServerInterceptor
}

// ServerConfig is the bm server config model
type ServerConfig struct {
	Network      string         `dsn:"network"`
	Addr         string         `dsn:"address"`
	Timeout      xtime.Duration `dsn:"query.timeout"`
	ReadTimeout  xtime.Duration `dsn:"query.readTimeout"`
	WriteTimeout xtime.Duration `dsn:"query.writeTimeout"`
}

func New2(svc pb.DemoServer) (h *Server, err error) {
	var (
		cfg ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	h = NewServer(&cfg)
	h, err = h.Start()
	return
}

func RegisterBMDemoServer() {}

func NewServer(conf *ServerConfig) (s *Server) {
	s = new(Server)
	s.SetConfig(conf)
	s.engine = gin.New()
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
		conf.Addr = "0.0.0.0:8000"
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
	return s, nil
}
