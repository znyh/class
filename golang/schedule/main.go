package main

import (
	"fmt"
	"sync"
	"time"
)

//gpm  内核线程与用户线程 n:m 的关系
//g：线程
//p：上下文调度器(cpu核数)
//m：执行线程的内核单元

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Wait()
		}(i)
	}

	wg.Done()
	time.Sleep(time.Second)
	fmt.Println("done")
}
