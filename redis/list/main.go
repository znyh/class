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

	mylist()
}

func mylist() {
	_, err := c.Do("LPUSH", "list1", "ele1", "ele2", "ele3")
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.String(c.Do("LPOP", "list1"))
	if err != nil {
		fmt.Println("redis POP error:", err)
	} else {
		fmt.Printf("res type : %s \n", reflect.TypeOf(res))
		fmt.Printf("%s \n", res)
	}

	datas, err := redis.Values(c.Do("LRANGE", "list1", "0", "-1"))
	for _, data := range datas {
		fmt.Printf("%s ", string(data.([]byte)))
	}

}
