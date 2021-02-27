package patterns

import (
	"testing"
)

func TestIterator(t *testing.T) {

	it := new(iterator)

	ch := genvistor(5)
	for v := range ch {
		it.put(v)
	}

	it.doRange()

	//fmt.Println(vt)
	//it.del(vt)
	//it.doRange()

}
