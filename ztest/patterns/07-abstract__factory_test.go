package patterns

import (
	"testing"
)

func TestAbstratorFactory(t *testing.T) {

	f := new(concreatefactory)
	p := f.produce().(*concreateproduce)

	if p.name != "kd" {
		t.Error("bad abstrator factory")
	}
}
