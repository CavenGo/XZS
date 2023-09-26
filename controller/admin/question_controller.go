package admin

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xzs/common"
	"xzs/model/request"
	"xzs/pkg/validatorutil"
	"xzs/service/admin"
)

type Question struct {
}

var QuestionApi = new(Question)

func (q *Question) PageList(c *gin.Context) {
	var req request.QuestionPageRequestVM
	err := c.ShouldBind(&req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	err = validatorutil.MyValidate(req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res := admin.QuestionPageListService(req)
	common.ResponseOk(c, res)
}

func (q *Question) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	err = admin.QuestionDeleteService(id)
	if err != nil {
		common.ResponseFailWithMsg(c, "删除失败")
		return
	}
	common.ResponseOk(c, nil)
}

func (q *Question) Select(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res := admin.QuestionSelectService(id)
	common.ResponseOk(c, res)
}

func (q *Question) Edit(c *gin.Context) {
	var req request.QuestionEditRequest
	err := c.ShouldBind(&req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	userName := c.GetString("userName")
	admin.QuestionEditService(req, userName)
	common.ResponseOk(c, nil)
}
