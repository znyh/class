package patterns

import (
	"fmt"
)

type object struct {
}

func (o *object) do() {
	fmt.Println(&o)
}

func genobjectpool(count int) chan *object {
	ch := make(chan *object, count)
	defer close(ch)

	for i := 0; i < count; i++ {
		ch <- new(object)
	}
	return ch
}
