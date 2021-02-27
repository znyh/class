package main

import (
	"fmt"
	"sync"
	"time"
)

//同步锁

func main() {
	m := make(map[int]int)
	mu := new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			m[i] = i
		}(i)

		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			if v, ok := m[i]; ok {
				fmt.Println("read->", i, v)
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(m)
}
