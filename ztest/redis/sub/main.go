package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	go sub()
	go push("hello world")
	time.Sleep(time.Second)
}

func sub() {
	c, _ := redis.Dial("tcp", ":6379", redis.DialDatabase(1))

	psc := redis.PubSubConn{c}
	psc.Subscribe("channel1")

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s, %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

func push(message string) {
	c, _ := redis.Dial("tcp", ":6379", redis.DialDatabase(1))

	_, err := c.Do("PUBLISH", "channel1", message)
	if err != nil {
		fmt.Println("push err:", err)
		return
	}

}
