package main

import (
	"container/list"
	"fmt"
	//"github.com/gin-gonic/gin"
)

type student struct {
	id   int
	name string
}

func main() {
	l := list.New()
	l.PushBack(11)
	l.PushBack("aaa")
	l.PushBack(student{id: 1001, name: "zyh"})

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) { c.String(200, "message") })
	// r.Run(":8080")

}
