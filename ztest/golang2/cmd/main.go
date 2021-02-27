package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

var (
	redisConn redis.Conn

	redisDB   = flag.Int("redisDB", 1, "redis db")
	redisAddr = flag.String("redisAddr", "localhost:6379", "redis listen addr")
)

type (
	StCreateTable struct {
		UserId   int32 `json:"user_id"`
		GameId   int32 `json:"game_id"`
		IsCreate bool  `json:"is_create"`
	}

	StCanCreateTable struct {
		UserId int32 `json:"user_id"`
		GameId int32 `json:"game_id"`
	}
)

func main() {

	initRedis()

	r := gin.Default()

	r.POST("/behavior/CreateTable", createTable)
	r.POST("/behavior/CanCreateTable", canCreateTable)

	r.Run(":9001")
}

func createTable(c *gin.Context) {

	var ct StCreateTable
	if err := c.ShouldBind(&ct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUserCreateNum(ct.GameId, ct.UserId, ct.IsCreate)

	c.JSON(200, gin.H{
		"user_id": ct.UserId,
		"game_id": ct.GameId,
	})
}

func canCreateTable(c *gin.Context) {

	var cct StCanCreateTable

	if err := c.ShouldBind(&cct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	can := userCanCreateTable(cct.GameId, cct.UserId)

	c.JSON(200, gin.H{
		"user_id":    cct.UserId,
		"game_id":    cct.GameId,
		"can_create": can,
	})
}

func initRedis() {
	flag.Parse()

	var err error
	if redisConn, err = redis.Dial("tcp", *redisAddr); err != nil {
		fmt.Println("redis conn err:", err)
		return
	}

	if _, err = redisConn.Do("SELECT", *redisDB); err != nil {
		redisConn.Close()
		return
	}
}

func updateUserCreateNum(gameId int32, userId int32, isCreate bool) {

	cmd := "INCR"
	if !isCreate {
		cmd = "DECR"
	}

	key := fmt.Sprintf("%d_%d", gameId, userId)

	_, err := redis.Int64(redisConn.Do(cmd, key))
	if err != nil {
		return
	}
}

func userCanCreateTable(gameId int32, userId int32) bool {
	key := fmt.Sprintf("%d_%d", gameId, userId)

	data, err := redis.String(redisConn.Do("GET", key))
	if err != nil || data == "" {
		redis.String(redisConn.Do("SET", key, "1"))
		return true
	} else {
		intN, _ := strconv.Atoi(data)
		return intN < 10
	}
}
