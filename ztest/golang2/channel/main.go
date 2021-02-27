package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan int, 10)
	go producer(2, ch)
	go producer(3, ch)
	go consumer(ch)
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func producer(x int, out chan<- int) {
	var i = 0
	for {
		out <- i * x
		i++
	}
}

func consumer(in <-chan int) {
	for i := range in {
		log.Printf("i = %d", i)
	}
}
