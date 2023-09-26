package admin

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xzs/common"
	"xzs/model/request"
	"xzs/model/response"
	"xzs/pkg/validatorutil"
	"xzs/service/admin"
)

type Task struct {
}

var TaskApi = new(Task)

func (t *Task) PageList(c *gin.Context) {
	var req request.TaskPageListRequest
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
	res := admin.TaskPageListService(req)
	common.ResponseOk(c, res)
}

func (t *Task) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	err = admin.TaskDeleteService(id)
	if err != nil {
		common.ResponseFailWithMsg(c, err.Error())
		return
	}
	common.ResponseOk(c, nil)
}

func (t *Task) Select(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res := admin.TaskSelectService(id)
	common.ResponseOk(c, res)
}

func (t *Task) Edit(c *gin.Context) {
	var req response.TaskRequestVM
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
	taskExam := admin.TaskEditService(req, userName)
	res := admin.TaskSelectService(taskExam.Id)
	common.ResponseOk(c, res)
}
