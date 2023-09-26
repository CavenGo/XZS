package admin

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xzs/common"
	"xzs/model/request"
	"xzs/pkg/validatorutil"
	"xzs/service/admin"
)

type ExamPaper struct {
}

var ExamPaperApi = new(ExamPaper)

func (e *ExamPaper) Page(c *gin.Context) {
	var req request.ExamPaperPageRequest
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
	res := admin.ExamPaperTaskExamPageListService(req)
	common.ResponseOk(c, res)
}

func (e *ExamPaper) TaskExamPageList(c *gin.Context) {
	var req request.ExamPaperPageRequest
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
	res := admin.ExamPaperTaskExamPageListService(req)
	common.ResponseOk(c, res)
}

func (e *ExamPaper) Edit(c *gin.Context) {
	var req request.ExamPaperEditRequest
	err := c.ShouldBind(&req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	err = validatorutil.MyValidate(req)
	if err != nil {
		common.ResponseFailWithMsg(c, err.Error())
		return
	}
	userName := c.GetString("userName")
	examPaper, err := admin.ExamPaperEditService(req, userName)
	if err != nil {
		common.ResponseFailWithMsg(c, err.Error())
		return
	}
	res, err := admin.ExamPaperSelectService(examPaper.Id)
	if err != nil {
		common.ResponseFailWithCode(c, common.InnerError)
		return
	}
	common.ResponseOk(c, res)
}

func (e *ExamPaper) Select(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res, err := admin.ExamPaperSelectService(id)
	if err != nil {
		common.ResponseFailWithMsg(c, err.Error())
		return
	}
	common.ResponseOk(c, res)
}

func (e *ExamPaper) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	err = admin.ExamPaperDeleteService(id)
	if err != nil {
		common.ResponseFailWithMsg(c, "删除失败")
		return
	}
	common.ResponseOk(c, nil)
}
