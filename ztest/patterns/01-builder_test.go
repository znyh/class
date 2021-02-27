package patterns

import (
	"testing"
)

func TestBuilder(t *testing.T) {

	c := new(car)

	d := new(director)
	d.construct(c)

	v := d.getvehicle()

	if v.seat != 4 || v.wheel != 4 {
		t.Error("bad builder")
	}
}
