package main

import (
	"fmt"
	"sync"
	"time"
)

//同步包
func main() {
	m := new(sync.Map)

	for i := 0; i < 5; i++ {
		go func(i int) {
			m.Store(i, i)
			fmt.Println("wirte ", i)
		}(i)

		go func(i int) {
			v, _ := m.Load(i)
			fmt.Println("read->", i, v)
		}(i)
	}
	time.Sleep(time.Second)
}
