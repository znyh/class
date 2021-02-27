package decorator

import (
	"fmt"
)

/*
	装饰模式: 使用对象组合的方式动态改变或增加对象行为， 在原对象的基础上增加功能(实现接口，扩展方法功能)
*/

type Component interface {
	GetCount() int
	Describe() string
}

//基础结构体
type Fruit struct {
	Count       int
	Description string
}

func (f *Fruit) GetCount() int {
	return f.Count
}

func (f *Fruit) Describe() string {
	return f.Description
}

//装饰结构体
type AppleDecorator struct {
	Component
	Type string
	Num  int
}

func (apple *AppleDecorator) GetCount() int {
	return apple.Component.GetCount() + apple.Num
}

func (apple *AppleDecorator) Describe() string {
	return fmt.Sprintf("%s, %s", apple.Component.Describe(), apple.Type)
}

func CreateAppleDecorator(c Component, t string, n int) Component {
	return &AppleDecorator{c, t, n}
}
