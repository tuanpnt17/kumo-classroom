package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanpnt17/kumo-classroom-be/internal/domain/models"
)

type AuthController struct {
	loginUsecase models.ILoginUsecase
}

func NewAuthController(loginUsecase models.ILoginUsecase) *AuthController {
	return &AuthController{
		loginUsecase: loginUsecase,
	}
}

func (ac *AuthController) Login(c *gin.Context) {

}

func (ac *AuthController) Register(c *gin.Context) {

}

func (ac *AuthController) Refresh(c *gin.Context) {

}
