package main

import (
	"flag"
	"fmt"
	"reflect"

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
	flag.Parse()

	initRedis(*redisAddr)

	c.Do("FLUSHDB")

	myhash()
}

func myhash() {
	_, err := c.Do("HSET", "student", "name", "wd", "age", 22)
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.Int64(c.Do("HGET", "student", "age"))
	if err != nil {
		fmt.Println("redis HGET error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("res  : %d \n", res)
	}
}
