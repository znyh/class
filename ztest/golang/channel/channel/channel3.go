package main

import (
	"fmt"
)

func merge(inputs ...<-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	for _, in := range inputs {
		in := in
		go func(in <-chan uint64) {
			for {
				output <- <-in
			}
		}(in)
	}
	return output
}

func divisor(input <-chan uint64, outputs ...chan<- uint64) {
	for _, out := range outputs {
		out := out
		go func(out chan<- uint64) {
			for {
				out <- <-input
			}
		}(out)
	}
}

func producer(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func square(inch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range inch {
			out <- n * n
		}
	}()

	return out
}

func main() {
	in := producer(1, 2, 3, 4)
	out := square(in)

	for n := range out {
		fmt.Println("===>", n)
	}
}
