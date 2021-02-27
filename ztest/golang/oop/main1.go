package main

import (
	"fmt"
)

type (
	People  struct{}
	Teacher struct {
		People
	}
)

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

//showA
//showB
//teacher showB
func main() {
	t := Teacher{}
	t.ShowA()
	t.ShowB()
}
