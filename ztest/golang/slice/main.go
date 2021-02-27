package main

import (
	"fmt"
)

//slice 的长度发生改变了意味着slice在内存中的地址也发生了改变
func main() {

	a := []string{"a", "b", "c"}
	b := a[1:2]
	fmt.Println(a, b)

	b[0] = "bbbbb"
	fmt.Println(a, b)

	a[1] = "nnnnn"
	fmt.Println(a, b)

	a = []string(nil)
	// a = []string{}
	fmt.Println(a, b)
}
