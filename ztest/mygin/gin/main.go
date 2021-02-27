package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const addr = ":9876"

func main() {

	//f, err := os.OpenFile("./gin.log", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 755)
	//if err != nil {
	//	return
	//}
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//defer f.Close()

	//r := gin.Default()

	r := gin.New()

	//r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//
	//	// 你的自定义格式
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))

	//r.Use(Logger())

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/ping", Ping)
	r.GET("/hello", Hello)
	r.GET("/param/:name", Param)
	r.GET("/query", Query)
	r.POST("/postform", PostForm)
	r.POST("/getpost", GetPost) //get+post混合

	g := r.Group("/group")
	{
		g.GET("/get", GroupTest)
		g.POST("/post", GroupTest)
	}

	//若要将请求主体绑定到结构体中，请使用模型绑定，目前支持JSON、XML、YAML和标准表单值(foo=bar&boo=baz)的绑定
	g2 := r.Group("/bind")
	{
		g2.POST("/binding", Binding)    //绑定Get参数或者Post参数
		g2.GET("/binding2", BindingGet) //只绑定Get参数
		g2.POST("/json", BindJson)
		g2.POST("/binding3", BindingBody) //绑定请求体
	}

	r.GET("/middle", Middle)

	log.Fatal(r.Run(addr))
}

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func Hello(c *gin.Context) {
	c.String(200, "hello world")
}

// curl -X GET "http://127.0.0.1:9876/param/xxx"
func Param(c *gin.Context) {
	name := c.Param("name")
	c.String(200, name)
}

// curl -X GET "http://127.0.0.1:9876/param/xxx"
func Query(c *gin.Context) {
	first := c.Query("first")
	second := c.DefaultQuery("second", "mmm")
	s := "hello, " + first + second
	c.String(200, s)
}

// curl -X POST "http://127.0.0.1:9876/postform" -d "first=zhang&second=san"
func PostForm(c *gin.Context) {
	first := c.PostForm("first")
	second := c.DefaultPostForm("second", "nnn")
	s := "hello, " + first + second
	c.String(200, s)
}

// curl -H "Content-Type:application/x-www-form-urlencoded" -X  POST "http://127.0.0.1:9876/getpost" -d "first=zhang&second=san"
func GetPost(c *gin.Context) {
	first := c.DefaultQuery("first", "xfirst")
	second := c.DefaultPostForm("second", "xsecond")

	str := first + "-" + second
	c.String(200, str)
}

// curl  -X  POST "http://127.0.0.1:9876/group/post"
func GroupTest(c *gin.Context) {
	c.String(200, "group test")
}

// url绑定参数（也可以绑定到body里）
type person struct {
	Name string `form:"name" binding:"required"`
	Id   int32  `form:"id" binding:"required"`
}

// body绑定struct
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// curl -X POST "http://127.0.0.1:9876/bind/binding" -d "name=zhangsan&id=1001"
func Binding(c *gin.Context) {
	var p person
	if err := c.ShouldBind(&p); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"name": p.Name,
		"id":   p.Id,
	})
}

// curl -X GET "http://127.0.0.1:9876/bind/binding2?name=zhangsan&id=1001"
func BindingGet(c *gin.Context) {
	var p person
	if err := c.ShouldBindQuery(&p); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"name": p.Name,
		"id":   p.Id,
	})
}

// curl  -X  POST "http://127.0.0.1:9876/group/post"
func BindJson(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "zhangsan" || json.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func Middle(c *gin.Context) {
	example := c.MustGet("example").(string)

	// it would print: "12345"
	log.Printf("example:%s", example)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func BindingBody(c *gin.Context) {
	var p person
	if err := c.ShouldBindBodyWith(&p, binding.JSON); err != nil {
		log.Printf("shouldBindBodyWith err:%+v", err)
	}
	c.JSON(200, p)
}
