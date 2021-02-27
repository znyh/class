package main

import (
	"log"
	"time"
)

func worker(id int, ready <-chan struct{}, done chan<- struct{}) {
	<-ready

	log.Println("worker:", id, " start to process.")

	time.Sleep(time.Second * 1)

	log.Println("worker:", id, " finish its job.")

	done <- struct{}{}
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan struct{}), make(chan struct{})

	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	ready <- struct{}{}
	ready <- struct{}{}
	ready <- struct{}{}

	<-done
	<-done
	<-done

}
