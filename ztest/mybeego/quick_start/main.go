package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Ctx.WriteString("hello beego")
}

func main() {
	//注册路由
	beego.Router("/", &HomeController{})
	beego.Run("127.0.0.1:8000")
}
