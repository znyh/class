package patterns

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {

	o := &prototype{desc: "init origin"}

	c := o.clone()

	if c.desc != o.desc {
		t.Error("bad prototype")
	}

	o.desc = "abc"
	fmt.Printf("%v,%v\n", o, c)

	c.desc = "123"
	fmt.Printf("%v,%v\n", o, c)
}
