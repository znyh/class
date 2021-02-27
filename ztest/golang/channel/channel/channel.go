package main

import (
	"fmt"
)

//阻塞？没准备好就往channel中取数据或放数据
//死锁? 所有goroutine都在等（阻塞），所以死锁。

//chan是一个FIFO队列，chan分成两种类型同步和异步
//同步的chan完成发送者和接受者之间手递手传递元素的过程，必须要求对方的存在才能完成一次发送或接受
//异步的chan发送和接受都是基于chan的缓存，但当缓存队列填满后，发送者就会进入发送队列, 当缓存队列为空时，接受者就会接入等待队列。

//缓冲channel：缓冲信道不仅可以流通数据，还可以缓存数据
//缓冲channel死锁的原因是因为channel满了或者空了

func main_() {
	c, quit := make(chan int), make(chan int)
	go func() {
		quit <- 0
		c <- 1 //不死锁的原因是main线程退出了
	}()
	<-quit // quit 等待数据的写
}

func main0() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 //发生死锁，容不下第4个数据
}

//缓冲信道看作为一个线程安全的队列
func main1() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3
}

//使用for来读取缓冲channel
func main2() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	for v := range ch {
		fmt.Println(v) //出现死锁
	}
}

func main3() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	//被关闭的信道会禁止数据流入, 是只读的。我们仍然可以从关闭的信道中取出数据，但是不能再写入数据了。
	close(ch) // 显式地关闭信道,避免死锁

	for v := range ch {
		fmt.Println(v)
	}
}
