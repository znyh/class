package main

import (
	"log"
	"math"
	"sync"
)

// 用于线程同步

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			sum(id)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}

	log.Println(id, x)
}
