package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./example.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open example.log err")

	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Fatal("f stat err")
	}
	fmt.Println(fi.Name(), fi.IsDir(), fi.Mode(), fi.ModTime(), fi.Size())

	//写入
	f.WriteString("我是一个好人\n")

	//读取
	if data, err := ioutil.ReadFile("./example.log"); err != nil {
		log.Fatal("ioutil readall fail")
	} else {
		fmt.Println(string(data))
	}

	//写入2
	// err = ioutil.WriteFile("./example.log", []byte("这是一个测试文件!!\n"), os.ModeAppend)
	// if err != nil {
	// 	log.Fatal("ioutil writeto file err")
	// }

	//读取2
	// if data, err := ioutil.ReadAll(f); err == nil {
	// 	fmt.Println(string(data))
	// } else {
	// 	log.Fatal("ioutil readall fail")
	// }
}
