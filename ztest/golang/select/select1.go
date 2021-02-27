package main

import (
	"fmt"
	"time"
)

//select 用于处理异步IO操作
//当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作

// 1.如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
// 2.如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
// 3.如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行

//用法1.超时判断
func main1() {
	var resChan = make(chan int)
	// do request
	for {
		select {
		case data := <-resChan:
			doData(data)
		case <-time.After(time.Second * 3):
			fmt.Println("request time out")
		}
	}
}
func doData(data int) {
	//...
}

//用法2.程序退出
func main2() {
	var shouldQuit = make(chan struct{})
	for {
		select {
		case data := <-shouldQuit:
			cleanUp(data)
			return
		default:
		}
	}
	//再另外一个协程中，如果运行遇到非法操作或不可处理的错误，就向shouldQuit发送数据通知程序停止运行
	close(shouldQuit)
}
func cleanUp(data struct{}) {
	//...
}

//用法3.判断channel是否阻塞
func main3() {
	//在某些情况下是存在不希望channel缓存满了的情况

	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full !")
	}
}
