package controller

import (
	"go-backend-api/internal/services"
	"go-backend-api/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService     *services.UserService
	responseService *response.ResponseData
}

func NewUserController() *UserController {
	return &UserController{
		userService:     services.NewUserService(),
		responseService: &response.ResponseData{},
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	var id = 2
	if id == 1 {
		response.SuccessResponse(c, 1)

	} else {
		response.ErrorResponse(c, 401)
	}
}
