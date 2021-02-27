package patterns

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	ch := generator(5)

	for v := range ch {
		fmt.Println(v)
	}

}
