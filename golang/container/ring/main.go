package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < r.Len(); i++ {
		r.Value = i
		fmt.Print(r.Value, " ")
		r = r.Next()
	}

	r.Do(func(p interface{}) {
		fmt.Print("#", p)
	})
}
