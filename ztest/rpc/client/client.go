package main

import (
	"context"
	"log"

	pb "github.com/znyh/class/ztest/rpc/api"

	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	req := &pb.HelloRequest{
		Name:                 "abc",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	c := pb.NewHelloDemoClient(conn)
	rsp, err := c.SayHello(context.Background(), req, )
	if err != nil {
		log.Printf("err:%+v", err)
	}
	log.Printf("rsp:%+v", rsp)
}
