package patterns

import (
	"testing"
)

func TestObjectpool(t *testing.T) {
	ch := genobjectpool(5)
	for v := range ch {
		v.do()
	}
}
