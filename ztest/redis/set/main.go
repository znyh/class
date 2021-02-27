package main

import (
	"flag"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	redisDB   = flag.Int("redisDB", 1, "redis db")
	redisAddr = flag.String("redisAddr", "localhost:6379", "redis listen addr")

	c redis.Conn
)

func initRedis(addr string) {
	var err error
	c, err = redis.Dial("tcp", addr)
	if err != nil {
		fmt.Println("redis conn err:", err)
		return
	}
}

func main() {

}
