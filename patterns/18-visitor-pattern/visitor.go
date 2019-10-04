package visitor

import "fmt"

/*
	访问者模式: 将对象的数据和操作分离, 传入不同的visitor，输出不同的行为

	设计思路：
		1. Visitor interface
		2. ConcreteVisitor struct
		3. Element interface(accept Visitor)
		4. ConcreteElement struct
		5. ElementContainer（包含Element list）非必须
*/

type IVisitor interface {
	Visit()
}

type ConcreteVisitorA struct {
	Name string
}

func (conV *ConcreteVisitorA) Visit() {
	fmt.Println("this is visitor A")
}

type ConcreteVisitorB struct {
	Name string
}

func (conV *ConcreteVisitorB) Visit() {
	fmt.Println("this is visitor B")
}

type IElement interface {
	Accept(visitor IVisitor)
}

type Element struct{}

func (e *Element) Accept(visitor IVisitor) {
	visitor.Visit()
}

type ElementContainer struct {
	list []IElement
}

func (container *ElementContainer) Add(element IElement) {
	if container == nil || element == nil {
		return
	}
	container.list = append(container.list, element)
}

func (container *ElementContainer) Delete(element IElement) {
	if len(container.list) <= 0 {
		return
	}
	for i, val := range container.list {
		if val == element {
			container.list = append(container.list[:i], container.list[i+1:]...)
			break
		}
	}
}
