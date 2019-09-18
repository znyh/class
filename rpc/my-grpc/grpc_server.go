package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pbArith "class/rpc/my-grpc/proto"

	"google.golang.org/grpc"
)

type arith int

func (c *arith) SayHello(ctx context.Context, req *pbArith.ArithReq) (*pbArith.ArithRsp, error) {
	rsp := &pbArith.ArithRsp{
		Id:  req.Id,
		Msg: "say hello",
	}

	return rsp, nil
}

func main() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rpc服务器开始监听：", ":9999")
	s := grpc.NewServer()
	pbArith.RegisterArithServer(s, new(arith))
	log.Fatal(s.Serve(l))
}
