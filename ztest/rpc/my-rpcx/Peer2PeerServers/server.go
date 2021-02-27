package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/znyh/class/ztest/rpc/my-rpcx/pb"

	"github.com/smallnest/rpcx/server"
)

var (
	addr       = flag.String("addr", "localhost:9991", "server address")
	consulAddr = flag.String("consulAddr", "localhost:8500", "consul address")
	basePath   = flag.String("util", "/cy_im", "consul prefix path")
)

type Arith int

func (this *Arith) Calu(ctx context.Context, req *pb.ArithRequest, rsp *pb.ArithResponse) error {
	rsp.C = req.A * req.B
	return nil
}

func main() {
	flag.Parse()

	fmt.Println("rpc server Listen at:", *addr)
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)
}
