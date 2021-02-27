package main

import (
	"fmt"
)

func main() {

	ch, quit := make(chan int), make(chan bool)

	go func() {
		x, y := 0, 1
		for i := 0; i < 10; i++ {
			x, y = y, x+y
			ch <- x
		}
		quit <- true
	}()

	for {
		select {
		case <-quit:
			fmt.Println("quit")
			return
		case data := <-ch:
			fmt.Print(data, " ")
		}
	}
}
