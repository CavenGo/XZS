package admin

import (
	"github.com/gin-gonic/gin"
	"xzs/common"
	"xzs/model/request"
	"xzs/pkg/validatorutil"
	"xzs/service/admin"
)

type Answer struct {
}

var AnswerApi = new(Answer)

func (a *Answer) PageJudgeList(ctx *gin.Context) {
	var req request.ExamPaperAnswerPageRequestVM
	err := ctx.ShouldBind(&req)
	if err != nil {
		common.ResponseFailWithCode(ctx, common.ParameterValidError)
		return
	}
	err = validatorutil.MyValidate(req)
	if err != nil {
		common.ResponseFailWithMsg(ctx, err.Error())
		return
	}
	res := admin.AnswerPageJudgeListService(req)
	common.ResponseOk(ctx, res)
}
