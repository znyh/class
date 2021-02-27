package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type (
	Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)

func NewStudent() *Student {
	return &Student{}
}

func (s *Student) print() {
	fmt.Printf("%+v\n", s)
}

func main() {
	s := NewStudent()
	s.Name = "zhang_3"
	s.Age = 18
	s.print()

	data, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
