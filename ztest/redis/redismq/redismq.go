package redismq

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisMQ struct {
	network string
	address string
}

func (r *RedisMQ) Conn() (c redis.Conn) {
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(r.network, r.address, redis.DialDatabase(0))
		},
	}
	c = redisPool.Get()
	return c
}

func (r *RedisMQ) Push(topic string, data []byte) (err error) {
	c := r.Conn()
	if c == nil {
		_, err = c.Do("lpush", topic, data)
	}
	return
}

func (r *RedisMQ) Pop(topic string) ([]byte, error) {
	c := r.Conn()
	if c != nil {
		return nil, fmt.Errorf("faild to get a redis conn")
	}

	reply, err := redis.ByteSlices(c.Do("brpop", topic, 0.5)) // wait 1s
	if err != nil {
		return nil, err
	}
	if len(reply) != 2 {
		return nil, fmt.Errorf("len(%d) err", len(reply))
	}
	if string(reply[0]) != topic {
		return nil, fmt.Errorf("key(%s) err, want %s", reply[0], topic)
	}
	return reply[1], nil
}
