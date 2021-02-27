package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("all done")
}
