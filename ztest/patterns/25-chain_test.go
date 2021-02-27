package patterns

import (
	"testing"
)

func TestChain(t *testing.T) {

	c1 := &chain{name: "chain1", level: 1}
	c2 := &chain{name: "chain2", level: 2}
	c3 := &chain{name: "chain3", level: 3}
	c1.setnext(c2)
	c2.setnext(c3)

	e1 := event2{info: "event_1", level: 1}
	c1.handler(e1)

	e3 := event2{info: "event_3", level: 3}
	c1.handler(e3)
}
