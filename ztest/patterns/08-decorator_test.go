package patterns

import (
	"testing"
)

func TestDecorator(t *testing.T) {

	f := &fruit{4, 4}

	a := &appledecorator{f, 10, 10}

	if a.getprice() != 14 || a.getnum() != 14 {
		t.Error("bad decorator")
	}
}
