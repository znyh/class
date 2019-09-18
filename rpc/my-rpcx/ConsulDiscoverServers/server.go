package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"class/rpc/my-rpcx/pb"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr       = flag.String("addr", "localhost:9991", "server address")
	consulAddr = flag.String("consulAddr", "localhost:8500", "consul address")
	basePath   = flag.String("base", "/rpcx_test", "prefix path")
)

type Arith struct {
	IpAddr string
}

func (this *Arith) Calu(ctx context.Context, req *pb.ArithRequest, rsp *pb.ArithResponse) error {
	rsp.C = req.A * req.B
	rsp.Addr = this.IpAddr
	return nil
}

func main() {
	flag.Parse()

	fmt.Println("rpc server 监听在:", *addr)
	s := server.NewServer()
	addRegistryPlugin(s)
	s.RegisterName("Arith", &Arith{IpAddr: *addr}, "")
	s.Serve("tcp", *addr)
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ConsulServers:  []string{*consulAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
