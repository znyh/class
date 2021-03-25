package main

import (
	"flag"

	"class/ztest/myconf/internel/config"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	config.LoadConfig()
	log.Info("main start")
	<-make(chan bool)
}
