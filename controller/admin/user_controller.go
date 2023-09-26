package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xzs/common"
	"xzs/model/request"
	"xzs/pkg/validatorutil"
	"xzs/service"
	"xzs/service/admin"
)

type User struct {
}

var UserApi = new(User)

func (u *User) PageList(c *gin.Context) {
	var req request.PageListRequest
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
	res := admin.PageListService(req)
	common.ResponseOk(c, res)
}

func (u *User) ChangeStatus(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	newStatus, err := admin.UserChangeStatus(id)
	if err != nil {
		common.ResponseFailWithMsg(c, err.Error())
		return
	}
	common.ResponseOk(c, newStatus)
}

func (u *User) Current(c *gin.Context) {
	userName := c.GetString("userName")
	res, err := admin.UserCurrentService(userName)
	if err != nil {
		common.ResponseFailWithMsg(c, err.Error())
		return
	}
	common.ResponseOk(c, res)
}

func (u *User) Select(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res, err := admin.UserSelectService(id)
	if err != nil {
		common.ResponseFailWithMsg(c, err.Error())
		return
	}
	common.ResponseOk(c, res)
}

func (u *User) Edit(c *gin.Context) {
	var req request.UserEditRequest
	err := c.Bind(&req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	err = validatorutil.MyValidate(req)
	if err != nil {
		common.ResponseFailWithCodeMsg(c, common.ParameterValidError, err.Error())
		return
	}
	res := admin.UserEditService(req)
	c.JSONP(http.StatusOK, res)
}

func (u *User) Update(c *gin.Context) {
	var req request.UserUpdateRequest
	err := c.Bind(&req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	userName := c.GetString("userName")
	err = service.UserUpdateService(userName, req)
	if err != nil {
		common.ResponseFailWithMsg(c, err.Error())
		return
	}
	common.ResponseOk(c, nil)

}

func (u *User) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	err = service.UserDeleteService(id)
	if err != nil {
		common.ResponseFailWithMsg(c, "删除失败")
		return
	}
	common.ResponseOk(c, nil)
}

func (u *User) SelectByUserName(c *gin.Context) {
	body := make([]byte, c.Request.ContentLength)
	_, err := c.Request.Body.Read(body)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	userName := string(body)
	if userName == "" {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res := admin.UserSelectByUserNameService(userName)
	common.ResponseOk(c, res)
}

func (u *User) EventPageList(c *gin.Context) {
	var req request.UserEventPageRequestVM
	err := c.ShouldBind(&req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res := admin.EventPageListService(req)
	common.ResponseOk(c, res)
}
