package main

import (
	"fmt"
	"time"
)

type (
	mytest struct {
		wake     chan bool
		exit     chan error
		interval time.Duration
		callback func() error
	}
)

func NewMytest(i time.Duration, cb func() error) *mytest {
	m := &mytest{
		wake:     make(chan bool, 1),
		exit:     make(chan error, 1),
		interval: i,
		callback: cb,
	}

	go m.start()

	return m
}

func (r *mytest) Close() {
	r.exit <- nil

	close(r.wake)
	close(r.exit)
}

func (r *mytest) Wake() {
	r.wake <- true
}

func (r *mytest) Wait() (err error) {
	return <-r.exit
}

func (r *mytest) start() {
	for {
		select {
		case <-r.exit:
			return
		case <-r.wake:
			fmt.Println("wake")
		case <-time.After(r.interval):
			fmt.Println("========>")
		}
		if err := r.callback(); err != nil {
			fmt.Println("chu xian err:", err.Error())
			return
		}
	}
}

func main() {
	//gr := routine.NewRoutine(2*time.Second, func() error { return nil })

	r := NewMytest(2*time.Second, func() error {
		fmt.Println("i am here")
		return nil
	})

	go func(r *mytest) {
		r.Wake()
	}(r)

	go func(r *mytest) {
		if tmp := r.Wait(); tmp == nil {
			fmt.Println("----@@@@@")
		} else {
			fmt.Println("----@@@@@---2")
		}
	}(r)

	select {}
}
