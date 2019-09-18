package main

import "fmt"

type (
	person interface {
		say()
	}

	man struct {
	}

	woman struct {
		person //显示实现接口
	}
)

func (m *man) say() {
	fmt.Println("man say !")
}

func (m *woman) say() {
	fmt.Println("woman say !")
}

func demo(p person) {
	p.say()
}

func main() {

	m := []interface{}{&man{}, &woman{}}

	for _, v := range m {
		if p, ok := v.(person); ok {
			demo(p)
		}
	}
}
