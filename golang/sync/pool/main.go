package main

import (
	"fmt"
	"sync"
)

func main() {
	p := new(sync.Pool)

	for i := 0; i < 5; i++ {
		p.Put(i)
		fmt.Println(p.Get())
	}
}
