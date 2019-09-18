package main

import (
	"fmt"
	"sync/atomic"
)

type buck struct {
	routines []chan int

	routinesNum uint64

	RoutineAmount uint64
}

func newBuck() *buck {
	b := new(buck)

	b.RoutineAmount = 3
	b.routinesNum = 0

	b.routines = make([]chan int, b.RoutineAmount)

	for i := uint64(0); i < b.RoutineAmount; i++ {
		c := make(chan int, 10)
		b.routines[i] = c
		go b.PushRoom(c)
	}

	return b
}

func (b *buck) PushRoom(c chan int) {
	fmt.Printf("创建pushroom\n")
	for {
		arg := <-c

		fmt.Println("arg:", arg)

		//if room = b.Room(arg.RoomId); room != nil {
		//	room.Push(&arg.P)
		//}

	}

}

func (b *buck) BroadcastRoom(arg int) {
	// 广播消息递增id
	num := atomic.AddUint64(&b.routinesNum, 1) % b.RoutineAmount
	fmt.Printf("BroadcastRoom RoomMsgArg :%d\n", arg)
	fmt.Printf("bucket routinesNum :%d\n", b.routinesNum)
	b.routines[num] <- arg

}

func main() {
	///newBuck()
	b := newBuck()
	b.BroadcastRoom(1234)

	b.BroadcastRoom(1235)
}
