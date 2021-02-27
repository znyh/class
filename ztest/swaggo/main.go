package main

import (
	"flag"

	"swaggo/api"
	"swaggo/cache/mongo"
	_ "swaggo/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	mgoURI = flag.String("mgo", "mongodb://localhost:27017/student", "mongo connection URI")
)

func main() {

	flag.Parse()

	mongo.MongdbInit(*mgoURI)

	r := gin.New()
	//r.GET("/loginJson/", api.LoginJson)
	r.GET("/userinfo/:id", api.Userinfo)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
