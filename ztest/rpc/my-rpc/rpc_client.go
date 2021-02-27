package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/znyh/class/ztest/rpc/my-rpc/pb"
)

func main() {
	conn, err := rpc.DialHTTP("tcp", ":9990")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := pb.ArithRequest{10, 5}
	var res pb.ArithResponse

	err = conn.Call("Arith.Calu", req, &res)
	if err != nil {
		log.Fatalln("Arith.Calu error: ", err)
	}
	fmt.Printf("req:%+v,rep:%+v", req, res)
}
