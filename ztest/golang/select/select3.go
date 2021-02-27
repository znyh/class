package main

import (
	"fmt"
	"runtime"
)

//select随机性
func main() {
	runtime.GOMAXPROCS(1)
	c, c2 := make(chan int, 1), make(chan string, 1)
	c <- 1
	c2 <- "hello"

	select {
	case value := <-c:
		fmt.Println(value)
	case value := <-c2:
		panic(value)
	}
}
