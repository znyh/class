package patterns

import (
	"testing"
)

func TestInterpreter(t *testing.T) {

	e := createExpress(1, content{val: "abc"}, content{val: "abc"})
	if e.interpretate() == false {
		t.Error("bad interpreter,equal")
	}

	c := createExpress(2, content{val: "abc"}, content{val: "c"})
	if c.interpretate() == false {
		t.Error("bad interpreter,contain")
	}
}
