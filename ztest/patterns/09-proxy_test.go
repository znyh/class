package patterns

import (
	"testing"
)

func TestProxy(t *testing.T) {
	p := new(actionproxy)
	p.doaction("run")
	p.doaction("read")
}
