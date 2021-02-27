package patterns

import (
	"testing"
)

func TestStrategy(t *testing.T) {

	add := new(addition)
	mulit := new(multiplication)

	o := new(operation)

	if ret := o.operate(add, 2, 3); ret != 5 {
		t.Error("bad strategy add")
	}

	if ret := o.operate(mulit, 2, 3); ret != 6 {
		t.Error("bad strategy mulit")
	}

}
