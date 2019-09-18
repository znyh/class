package main

import (
	"context"
	"flag"
	"fmt"

	"class/rpc/my-rpcx/pb"

	"github.com/smallnest/rpcx/server"
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
	var (
		addr1 = flag.String("addr1", "localhost:9991", "server1 address")
		addr2 = flag.String("addr2", "localhost:9992", "server2 address")
	)

	flag.Parse()
	go createServer(*addr1)
	go createServer(*addr2)
	select {}
}

func createServer(addr string) {
	fmt.Println("rpc server 监听在:", addr)
	s := server.NewServer()
	s.RegisterName("Arith", &Arith{IpAddr: addr}, "")
	s.Serve("tcp", addr)
}
