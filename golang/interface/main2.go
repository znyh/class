package main

import (
	"fmt"
)

type Sayer interface {
	Say(message string)
	SayHi()
}
type Animal struct {
	Sayer
	Name string
}

func (a *Animal) Say(message string) {
	fmt.Printf("Animal[%v] say: %v\n", a.Name, message)
}

func (a *Animal) SayHi() {
	fmt.Printf("Animal SayHi")
}

type Dog struct {
	Animal
}

func (a *Dog) Say(message string) {
	fmt.Printf("Dog say\n")
}

func main() {
	d := new(Dog)
	d.Say("hello")
	d.SayHi()
}
