package main

import (
	"fmt"
)

//关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数

func main() {
	tmp := DeferFunc1(1)
	tmp2 := DeferFunc2(1)
	tmp3 := DeferFunc3(1)
	fmt.Println(tmp)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
}

// 4
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

// 1
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

// 3
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
