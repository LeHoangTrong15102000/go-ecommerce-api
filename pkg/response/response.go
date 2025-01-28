package response

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // Dữ liệu return
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {}