package admin

import (
	"github.com/gin-gonic/gin"
	"xzs/common"
	"xzs/model/request"
	"xzs/pkg/validatorutil"
	"xzs/service/admin"
)

type Message struct {
}

var MessageApi = new(Message)

func (m *Message) Send(c *gin.Context) {
	var req request.MessageSendVM
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
	userName := c.GetString("userName")
	err = admin.MessageSend(req, userName)
	if err != nil {
		common.ResponseFailWithMsg(c, err.Error())
		return
	}
	common.ResponseOk(c, nil)
}

func (m *Message) PageList(ctx *gin.Context) {
	var req request.MessagePageRequestVM
	err := ctx.ShouldBind(&req)
	if err != nil {
		common.ResponseFailWithCode(ctx, common.ParameterValidError)
		return
	}
	res := admin.MessagePageListService(req)
	common.ResponseOk(ctx, res)
}
