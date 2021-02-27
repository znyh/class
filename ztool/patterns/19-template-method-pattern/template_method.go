package template

import "fmt"

/*
	模版方法: 父类定义公共方法，不同子类重写父类抽象方法，得到不同结果
*/

type Shape interface {
	SetName(name string)
	BeforeAction()
	Exit()
}

type Person struct {
	name     string
	Concrete Shape
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) BeforeAction() {
	p.Concrete.BeforeAction()
}

func (p *Person) Exit() {
	p.BeforeAction()
	fmt.Println(p.name + "exit")
}

//匿名组合实现继承
type Boy struct {
	Person
}

func (b *Boy) BeforeAction() {
	fmt.Println(b.name)
}
