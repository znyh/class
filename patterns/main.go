package main

import "fmt"

type P struct {
	Name string
}

func main() {
	q := make(chan int, 3)
	q <- 1
	q <- 2
	q <- 3
	close(q)
	for v := range q {
		fmt.Println(v)
	}
	fmt.Println("jie")

	s := make(chan struct{}, 1)
	select {
	case s <- struct{}{}:
		fmt.Println("struct")
	}

	type sem chan struct{}
	m := make(sem, 1)
	select {
	case m <- struct{}{}:
		fmt.Println("struct1")
	}

	ch := make(chan int, 1)
	select {
	case ch <- 1:
		fmt.Println("jie ")
	}
}
