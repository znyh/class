package main

import (
	"context"
	"flag"
	"fmt"

	"class/rpc/my-rpcx/pb"

	"github.com/smallnest/rpcx/client"
)

func main() {
	Peer2Many()
}

func Peer2Many() {

	var (
		addr1 = flag.String("addr1", "localhost:9991", "server1 address")
		addr2 = flag.String("addr2", "localhost:9992", "server2 address")
	)

	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
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
