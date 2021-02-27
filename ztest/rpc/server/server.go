package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/znyh/class/ztest/rpc/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	log.Printf("req:%+v rsp:%+v", in, resp)
	return resp, nil
}

func main1() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	//// TLS认证
	//creds, err := credentials.NewServerTLSFromFile(filepath.Base("../../configs/server.pem"), filepath.Base("../configs/server.key"))
	//if err != nil {
	//	grpclog.Fatalf("Failed to generate credentials %v", err)
	//}
	//// 实例化grpc Server, 并开启TLS认证
	//s := grpc.NewServer(grpc.Creds(creds))
	s := grpc.NewServer()
	// 注册HelloService
	pb.RegisterHelloDemoServer(s, HelloService)
	grpclog.Println("Listen on " + Address + " with TLS")
	s.Serve(listen)
}

func main() {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloDemoServer(s, &HelloService)
	log.Printf("Listen on " + Address)
	s.Serve(lis)
}

//func main() {
//	//s := grpc.NewServer()
//	//engine := bm.DefaultServer()
//	//pb.RegisterHelloDemoBMServer(engine, s)
//	//engine.Start()
//}

//// New new a tcp server.
//func New(svc pb.ChessCometServer) (cs *comet.Server, err error) {
//	var (
//		tc struct {
//			Server *comet.ServerConfig
//		}
//	)
//	if err = paladin.Get("tcp.toml").UnmarshalTOML(&tc); err != nil {
//		if err != paladin.ErrNotExist {
//			return
//		}
//		err = nil
//	}
//	rand.Seed(time.Now().UTC().UnixNano())
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	cs = comet.NewServer(tc.Server)
//	pb.RegisterChessCometServer(cs, svc)
//	err = cs.StartTCP(runtime.NumCPU())
//	if err != nil {
//		panic(err)
//	}
//	return
//}
