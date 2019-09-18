package main

import (
	"fmt"
)

type error interface {
	Error() (string, int)
}

type ApiError struct {
	s string
	c int
}

func (e *ApiError) Error() (string, int) {
	return e.s, e.c
}

func New(text string, code int) error {
	return &ApiError{s: text, c: code}
}

func main() {
	data := New("ss", 100)
	fmt.Println(data)
}
