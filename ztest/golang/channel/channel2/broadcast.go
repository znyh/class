package main

import (
	"fmt"
)

type room struct {
	size      int
	broadcast []chan int
}

func newRoom(size int) *room {
	r := new(room)
	r.size = size
	r.broadcast = make([]chan int, size)
	for i := 0; i < size; i++ {
		r.broadcast[i] = make(chan int, 0)
	}
	return r
}

func main() {
	r := newRoom(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			r.broadcast[i] <- i

		}(i)
		//fmt.Printf("线程i:%d,接收数据：%d\n", i, <-r.broadcast[i])
	}

	fmt.Println("main done")
}
