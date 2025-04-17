package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpnt17/kumo-classroom-be/internal/services"
	"github.com/tuanpnt17/kumo-classroom-be/pkg/response"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetAll(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		response.ErrorResponse(c, response.ParamInvalid)
		return
	}
	response.SuccessResponse(c, response.Success, uc.userService.GetAllUsers())
}

func (uc *UserController) GetUserById(context *gin.Context) {

}
