package redismq

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
)

var (
	_net   = "tcp"
	_addr  = "localhost:6379"
	_topic = "TestMQ"
)

type Message struct {
	Topic string
	data  []byte
}

func TestMq(t *testing.T) {

	mq := &RedisMQ{network: _net, address: _addr}

	go func() {
		i := 0
		for {
			data := []byte(strconv.Itoa(i))
			if err := mq.Push(_topic, data); err != nil {
				log.Error("produce element err: %+v", data)
			} else {
				log.Info("produce element:%+v", data)
			}
			i++
			i = i % 65536
			<-time.After(time.Second * 5)
		}
	}()

	go func() {
		for {
			data, err := mq.Pop(_topic)
			if err == nil && len(data) > 0 {
				log.Info("consumer element:%+v", string(data))
			}
			<-time.After(time.Second * 7)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("main exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
