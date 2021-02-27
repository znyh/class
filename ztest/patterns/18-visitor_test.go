package patterns

import (
	"testing"
)

func TestVistor(t *testing.T) {
	d := new(door)

	va := &visitor{"a"}
	vb := &visitor{"b"}
	vc := &visitor{"c"}

	d.accept(va, vb, vc)

	d.remove(va, vc)
}
