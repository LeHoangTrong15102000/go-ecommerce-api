package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {}


// Một func return về con trỏ của controller
func NewPongController() *PongController {
	// return về địa chỉ con trỏ của struct PongController
	return &PongController{}
}


func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "LeHoangTrong")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"uid":     uid,
		"users":   []string{"user1", "user2", "user3"},
	})
}