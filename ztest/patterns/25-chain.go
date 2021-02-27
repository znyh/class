package patterns

import (
	"fmt"
)

type ichain interface {
	setnext(ichain)
	handler(event2)
}

type event2 struct {
	info  string
	level int
}

type chain struct {
	next  ichain
	name  string
	level int
}

func (c *chain) setnext(next ichain) {
	c.next = next
}

func (c *chain) handler(e event2) {
	if c.level == e.level {
		fmt.Printf("%s处理事件%+v\n", c.name, e)
	} else if c.next != nil {
		c.next.handler(e)
	} else {
		fmt.Printf("无法处理事件%+v\n", e)
	}
}
