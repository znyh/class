package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mutex sync.Mutex
	m := make(map[int]int)

	for i := 1; i < 4; i++ {
		go func(i int) {
			//fmt.Printf("===> Lock. (G%d)\n", i)
			mutex.Lock()
			defer mutex.Unlock()
			m[i] = i * i
			fmt.Printf("===> Lock. (G%d):%d\n", i, m[i])

		}(i)
	}

	time.Sleep(time.Second)
}
