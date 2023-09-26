package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xzs/common"
	"xzs/model/request"
	"xzs/pkg/validatorutil"
	"xzs/service"
)

type User struct{}

var UserApi = new(User)

// Login 登录，前台与后台通用
func (u *User) Login(c *gin.Context) {
	var req request.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	err = validatorutil.MyValidate(req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res := service.LoginService(c, req)
	c.JSON(http.StatusOK, res)
}

// Logout 退出，前台与后台通用
func (u *User) Logout(c *gin.Context) {
	service.LogoutService(c)
	common.ResponseOk(c, nil)
}
