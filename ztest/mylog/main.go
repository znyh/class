package main

import (
	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("hello world!!")
}
