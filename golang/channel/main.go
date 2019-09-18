package main

import "fmt"

//channel用于协程间的通信

//无缓冲的channel
func main1() {
	c := make(chan int)

	////1> 主线程死锁，无缓冲区既不能取出 也不能存入
	//c <- 99
	//fmt.Println(<-c)

	////2> 主线程死锁，主线程阻塞，子线程不能执行
	// c <- 99
	// go func() {
	// 	fmt.Println(<-c)
	// }()

	// //3> 主线程阻塞，直到其他线程写入或读取缓冲
	// go func() {
	// 	fmt.Println(<-c)
	// }()
	// c <- 99
	// fmt.Println("main done")

	go func() {
		c <- 99
	}()
	fmt.Println(<-c)
}

//带缓冲的channel
func main() {
	c := make(chan int, 2)

	c <- 99

	////用法1
	//fmt.Println(<-c)

	////用法2
	for {
		if v, ok := <-c; ok {
			fmt.Println(v)
			close(c)
			break
		}
	}
}
