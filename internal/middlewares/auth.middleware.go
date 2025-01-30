package middlewares

import (
	"fmt"
	"go-ecommerce-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthMiddleware")
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken, "Token is invalid")
			c.Abort()
			return
		}
	}
}
