package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Pong)
	}
	// r.GET("/ping", Pong)

	return r
}


func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "LeHoangTrong")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"uid": uid,
		"users": []string{"user1", "user2", "user3"},
	})
}