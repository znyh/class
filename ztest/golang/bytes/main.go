package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

//https://studygolang.com/pkgdoc

func main() {
	str1 := "aabbbcccddddefgh"
	str2 := "aabb"
	fmt.Println(bytes.Contains([]byte(str1), []byte(str2)))

	main2()
}

func main2() {

	bs := []byte(nil)
	bf := bytes.NewBuffer(bs)

	f, err := os.OpenFile("./example.log", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = bf.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bf.String())
}
