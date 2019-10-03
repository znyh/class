package pool

import "fmt"

/*
	对象池模式: 用于对象的生成成本大于维持成本

	*设计思想
		1.对象结构体
		2.类型为结构体指针的channel
		3.New方法, 创建新的对象放到channel中
*/

type Object struct {
	Name string
}

func (obj *Object) Do() {
	fmt.Println(&obj)
}

type Pool chan *Object

func NewPool(count int) *Pool {
	pool := make(Pool, count)
	defer close(pool)

	for i := 0; i < count; i++ {
		pool <- new(Object)
	}
	return &pool
}
