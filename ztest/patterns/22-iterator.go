package patterns

import (
	"fmt"
	"sync"
)

type ivisit interface {
	visit()
	setnext(ivisit)
	getnext() ivisit
	setprev(ivisit)
	getprev() ivisit
}

type vistor struct {
	name string
	prev ivisit
	next ivisit
}

func genvistor(count int) chan ivisit {
	ch := make(chan ivisit)
	go func(ch chan ivisit) {
		for i := 0; i < count; i++ {
			ch <- &vistor{name: fmt.Sprintf("visitor_%d", i)}
		}
		close(ch)
	}(ch)
	return ch
}

func (v *vistor) visit() {
	fmt.Printf("%s accept visit\n", v.name)
}

func (v *vistor) setnext(iv ivisit) {
	v.next = iv
}
func (v *vistor) getnext() ivisit {
	return v.next
}

func (v *vistor) setprev(iv ivisit) {
	v.prev = iv
}
func (v *vistor) getprev() ivisit {
	return v.prev
}

type iterator struct {
	length int
	next   ivisit
	sync.RWMutex
}

func (it *iterator) put(v ivisit) {
	if v == nil {
		return
	}

	if it.next != nil {
		it.next.setprev(v)
	}
	v.setnext(it.next)
	v.setprev(nil)
	it.next = v
	it.length++
}

func (it *iterator) del(v ivisit) {
	if v == nil {
		return
	}

	it.RLock()
	defer it.RUnlock()

	if v.getnext() != nil {
		v.getnext().setprev(v.getprev())
	}

	if v.getprev() != nil {
		v.getprev().setnext(v.getnext())

	} else {
		it.next = v.getnext()
	}
	it.length--
}

func (it *iterator) doRange() {
	for curr := it.next; curr != nil; curr = curr.getnext() {
		curr.visit()
	}
}
