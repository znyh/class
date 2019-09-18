package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"class/rpc/my-rpc/pb"
)

type Arith int

func (this Arith) Calu(req pb.ArithRequest, res *pb.ArithResponse) error {
	res.Mul = req.A * req.B
	res.Add = req.A + req.B
	res.Sub = req.A - req.B
	return nil
}

func main() {
	rpc.Register(new(Arith)) // 注册rpc服务
	rpc.HandleHTTP()         // 采用http协议作为rpc载体

	lis, err := net.Listen("tcp", ":9990")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	fmt.Printf("start Listen at: %s\n", lis.Addr())

	http.Serve(lis, nil)
}
