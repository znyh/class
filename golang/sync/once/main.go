package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	once := new(sync.Once)
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			wg.Done()

			once.Do(func() {
				fmt.Println(i)
			})

		}(i)
	}
	wg.Wait()
	time.Sleep(1000)
	fmt.Println("done")
}
