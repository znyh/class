package main

import (
	"fmt"
)

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

// 0 1 zz zz 4
func main() {
	fmt.Println(x, y, z, k, p)
}
