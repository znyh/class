package main

import "fmt"

//defer在return语句之后执行
func main() {
	fmt.Println(func1())
	fmt.Println(func2())
	fmt.Println(func3())

	main2()
}

// 1，100，101
func func1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func func2() (r int) {
	m := 100
	defer func() {
		m += 100
	}()
	return m
}

func func3() (r int) {
	defer func() {
		r += 100
	}()
	return 1
}

//defer在声明的时候其变量已经定义好了
// c 2, b 1 , a 1
func main2() {
	i := 1
	defer fmt.Println("a", i)
	defer func(j int) {
		fmt.Println("b", j)
	}(i)
	defer func() {
		fmt.Println("c", i)
	}()
	i++
}
