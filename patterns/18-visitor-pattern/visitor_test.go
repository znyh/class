package visitor

import "testing"

func TestElementA_Accept(t *testing.T) {
	elementA := &Element{}
	visitorA := &ConcreteVisitorA{Name: "lee"}
	visitorB := &ConcreteVisitorB{Name: "anne"}
	elementA.Accept(visitorA)
	elementA.Accept(visitorB)
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
