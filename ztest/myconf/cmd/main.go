package main

import (
	"flag"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/znyh/class/ztest/myconf/internel/config"
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
