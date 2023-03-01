package controller

import (
	"career/helper"
	"career/model/web"
	"career/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (cl *AuthControllerImpl) Login(c *gin.Context) {
	req := web.LoginRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}
	res, err := cl.AuthService.Login(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
	})
}
