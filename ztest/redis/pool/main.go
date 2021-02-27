package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	redisPool *redis.Pool
)

func main() {
	redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379", redis.DialDatabase(0))
		},
	}

	c := redisPool.Get()
	c.Do("SET", "name", "zzyh")
	data, _ := redis.String(c.Do("GET", "name"))
	fmt.Println("data->", data)
}
