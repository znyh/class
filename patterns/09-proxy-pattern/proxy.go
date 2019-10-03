package proxy

import "fmt"

/*
 	代理模式用于延迟处理操作或者在进行实际操作前后对真实对象进行其它处理。

	//设计思想
		1. 代理inteface
		2. 真实对象Object struct
		3. 代理对象ProxyObject struct，属性为Object, 拦截所有的action
		4. 方法ObjDo处理所有的逻辑
*/

type IObject interface {
	ObjDo(action string)
}

type Object struct {
	action string
}

func (obj *Object) ObjDo(action string) {
	fmt.Printf("I can %s", action)
}

type ProxyObject struct {
	object *Object
}

//拦截作用
func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = new(Object)
	}
	if action == "run" {
		p.object.ObjDo(action)
	}
}
