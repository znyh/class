package main

import (
	"context"
	"fmt"
	"log"

	pbArith "class/rpc/my-grpc/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	req := pbArith.ArithReq{
		Id: 10001,
	}
	c := pbArith.NewArithClient(conn)
	rsp, err := c.SayHello(context.TODO(), &req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("接收到rpc server远程数据，rsp：%+v\n", rsp)
}
