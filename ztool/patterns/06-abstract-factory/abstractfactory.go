package abstractfactory

import (
	"fmt"
)

/*
	抽象工厂模式:提供一个创建一系列相关或相互依赖对象的接口, 而无需指定它们具体的类
	不同条件下创建不同实例
*/

type Factory interface {
	CreateProduct() Product
}

type Product interface {
	Describe()
}

//具体的产品
type ConcreteProduct struct {
	Name string
}

func (conproduct *ConcreteProduct) Describe() {
	fmt.Println(conproduct.Name)
}

//具体工厂
type ConCreteFactory struct{}

func (confactory *ConCreteFactory) CreateProduct() Product {
	return &ConcreteProduct{Name: "KG"}
}
