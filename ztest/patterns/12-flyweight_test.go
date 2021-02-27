package patterns

import (
	"testing"
)

func TestFlyweight(t *testing.T) {

	f := new(circleFactory)
	c := f.getcircle(5).(*circle)

	if c.color != 5 {
		t.Error("bad flyweight")
	}
}
