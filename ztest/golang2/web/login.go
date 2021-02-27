package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	stLoginJsonReq struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	stLoginJsonRsp struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Token   string `json:"token"`
	}
)

func main() {
	r := gin.Default()

	r.POST("/login", userLogin)

	r.Run(":9999")
}

func userLogin(c *gin.Context) {
	var login stLoginJsonReq
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	token := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s_%s", login.Name, login.Password))))
	c.JSON(http.StatusOK, stLoginJsonRsp{
		Code:    0,
		Message: "success",
		Token:   token,
	})

}
