package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpnt17/kumo-classroom-be/internal/services"
	"github.com/tuanpnt17/kumo-classroom-be/pkg/response"
)

type UserController struct {
	userService *services.IUserService
}

func NewUserController(us services.IUserService) *UserController {
	return &UserController{
		userService: &us,
	}
}

func (uc *UserController) GetAll(c *gin.Context) {
	response.SuccessResponse(c, response.Success, (*uc.userService).GetAllUsers())
}

func (uc *UserController) GetUserById(context *gin.Context) {

}
