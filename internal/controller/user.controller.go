package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)


type UserController struct {
	userService *service.UserService
}


// Một func return về con trỏ của controller
func NewUserController() *UserController {
	// return về địa chỉ con trỏ của struct UserController
	return &UserController{
		userService: service.NewUserService(),
	}
}


// uc user controller
// us user service đây là những quy ước chung của golang khi mà phát triển phần mềm
func (uc *UserController) GetUserById(c *gin.Context) {

	// response.SuccessResponse(c, 20001, []string{"user1", "user2", "user3"})
	response.ErrorResponse(c, 40001, "Error no need")
}