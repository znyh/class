package main

import (
	"flag"
	"fmt"
	"time"

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

	//mystring2()
	mystring3()
	mystring4()
}

func mystring1() {
	c.Do("SET", "mykey1", "zhang")
	c.Do("SET", "mykey2", "superHao")

	data1, _ := redis.String(c.Do("GET", "mykey1"))
	data2, _ := redis.String(c.Do("GET", "mykey2"))

	fmt.Println(data1, data2)
}

func mystring2() {
	c.Do("SET", "mykey3", "superHao", "EX", "3")
	data, _ := redis.String(c.Do("GET", "mykey3"))
	fmt.Println(data)

	data2, _ := redis.Bool(c.Do("EXISTS", "mykey3"))
	fmt.Println("exist mykey3:", data2)

	time.Sleep(time.Second * 5)

	data3, _ := redis.Bool(c.Do("EXISTS", "mykey3"))
	fmt.Println("exist mykey3:", data3)
}

func mystring3() {
	c.Do("MSET", "name", "zhang", "age", "18")
	res, _ := redis.String(c.Do("MGET", "name", "age"))
	fmt.Printf("%s", res)
}

func mystring4() {
	c.Do("SET", "mykey4", "zhang SuperHao")
	c.Do("EXPIRE", "mykey4", 10)
	data, _ := redis.String(c.Do("GET", "mykey4"))
	fmt.Println(data)
}
