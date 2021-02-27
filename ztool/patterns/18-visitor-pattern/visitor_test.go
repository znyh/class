package visitor

import "testing"

func TestElement_Accept(t *testing.T) {
	element := &Element{}
	visitorA := &ConcreteVisitorA{Name: "lee"}
	visitorB := &ConcreteVisitorB{Name: "anne"}
	element.Accept(visitorA)
	element.Accept(visitorB)
}

func TestElementContainer_Add(t *testing.T) {
	container := new(ElementContainer)
	a := &Element{}
	b := &Element{}
	container.Add(a)
	container.Add(b)
	if len(container.list) != 2 {
		t.Error("count error, expected amount is 2")
	}
}
