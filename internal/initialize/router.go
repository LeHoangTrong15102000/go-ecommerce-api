package initialize

import (
	"fmt"
	c "go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("Alter --> AA")
	}
}


func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("Alter --> BB")
	}
}


func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
		c.Next()
		fmt.Println("Alter --> CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()


	// use the middleware
	r.Use(middlewares.AuthMiddleware(), BB(), CC)

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
