package main

import (
	"fmt"
	"sync"
)

type (
	threadSafeSet struct {
		sync.RWMutex
		s []int
	}
)

func NewThreadSafeSet(s []int) *threadSafeSet {
	return &threadSafeSet{
		s: s,
	}
}
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
		for elem := range set.s {
			ch <- elem
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {
	t := NewThreadSafeSet([]int{1, 2, 3, 4})
	c := t.Iter()

	for v := range c {
		fmt.Println(v)
	}
}
