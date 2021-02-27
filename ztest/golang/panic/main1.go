package main

import (
	"fmt"
)

//defer 后进先出
//panic 需要等defer 结束后才会向上传递
//先按照defer的后入先出的顺序执行，最后才会执行panic。

func main() {

	defer func() { fmt.Println(1) }()
	defer func() { fmt.Println(2) }()
	defer func() { fmt.Println(3) }()

	panic("==> panic")
}
