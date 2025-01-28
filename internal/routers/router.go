package routers

import (
	c "go-ecommerce-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewPongController().Pong) // Còn để như này thì không phân biệt được gì hết
		v1.GET("/users/1", c.NewUserController().GetUserById) // Thì khi mà sửa lại thì sẽ viết ra thành như thế này
	}

	// v2 := r.Group("/v2/2025")
	// {
	// 	v2.GET("/ping", Pong)
	// 	v2.POST("/ping", Pong)
	// 	v2.PUT("/ping", Pong)
	// 	v2.PATCH("/ping", Pong)
	// }
	// r.GET("/ping", Pong)

	return r
}
