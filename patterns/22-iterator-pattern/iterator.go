package iterator

import "fmt"

/*
	迭代器模式:可以配合访问者模式，将不同的数据结构，使用迭代器遍历

	设计思想：
		1. Iterator结构体
			实现Next()  HasNext()方法
		2. Container容器
			容器实现添加 移除Visitor 和
*/

type Visitor interface {
	Visit()
}

type Teacher struct{}

func (t *Teacher) Visit() {
	fmt.Println("this is teacher visitor")
}

type Analysis struct{}

func (a *Analysis) Visit() {
	fmt.Println("this is analysis visitor")
}

type Container struct {
	list []Visitor
}

func (c *Container) Add(visitor Visitor) {
	c.list = append(c.list, visitor)
}

func (c *Container) Remove(index int) {
	if index < 0 || index > len(c.list) {
		return
	}
	c.list = append(c.list[:index], c.list[index+1:]...)
}

type Iterator struct {
	index int
	Container
}

func NewIterator() *Iterator {
	return &Iterator{
		index:     0,
		Container: Container{},
	}
}

func (i *Iterator) Next() Visitor {
	fmt.Println(i.index)
	visitor := i.list[i.index]
	i.index += 1
	return visitor
}

func (i *Iterator) HasNext() bool {
	if i.index >= len(i.list) {
		return false
	}
	return true
}
