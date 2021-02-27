package main

import "fmt"

//map是线程不安全的

type student struct {
	name string
}

func main() {
	ss := []student{student{name: "aaa"}, student{name: "bbb"}, student{name: "ccc"}}

	m := make(map[int]*student)

	for k, v := range ss {
		fmt.Println(&v) //&{"aaa"},&{"bbb"},&{"ccc"}
		m[k] = &v       //m[k] = &ss[k] ???
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
