package main

import (
	"context"
	"flag"
	"fmt"

	"class/rpc/my-rpcx/pb"

	"github.com/smallnest/rpcx/client"
)

func main() {

	var (
		consulAddr = flag.String("consulAddr", "localhost:8500", "consul address")
		basePath   = flag.String("util", "/rpcx_test/Arith", "prefix path")
	)

	d := client.NewConsulDiscovery(*basePath, "", []string{*consulAddr}, nil)
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	req := &pb.ArithRequest{A: 2, B: 5}
	rsp := &pb.ArithResponse{}

	err := xclient.Call(context.Background(), "Calu", req, rsp)
	if err != nil {
		fmt.Println("xclient Call err:", err)
		return
	}
	fmt.Printf("rpc rpc-client 远程调用，req:%+v,rsp:%+v\n", req, rsp)
}
