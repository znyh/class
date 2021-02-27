package api

import (
	"fmt"
	"net/http"
	"strconv"

	"swaggo/cache/mongo"
	"swaggo/web"

	"github.com/gin-gonic/gin"
)

type apiRsp struct {
	Code web.RspCode `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// @Summary 查询用户信息
// @Description
// @Accept  json
// @Produce json
// @Param   userid     path    int     true        "用户ID"
// @Success 0 {object} student.StudentInfo  "用户信息"
// @Router /userinfo/{userid} [get]
func Userinfo(c *gin.Context) {
	uidStr := c.Param("id")
	if uidStr == "" {
		return
	}

	uid, err := strconv.ParseUint(uidStr, 10, 64)
	fmt.Println("查询uid:", uid)
	if err != nil {
		c.JSON(http.StatusOK, apiRsp{Code: web.ArgInvalid, Msg: err.Error()})
		return
	}

	info, err := mongo.QueryStudentInfo(uid)
	if err != nil {
		c.JSON(http.StatusOK, apiRsp{Code: web.NotFound, Msg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, apiRsp{Code: web.Succ, Data: info})
}
