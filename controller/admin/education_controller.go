package admin

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xzs/common"
	"xzs/model/request"
	"xzs/pkg/validatorutil"
	"xzs/service/admin"
)

type Education struct {
}

var EducationApi = new(Education)

func (e *Education) List(c *gin.Context) {
	res := admin.EducationListService()
	common.ResponseOk(c, res)
}

func (e *Education) PageList(c *gin.Context) {
	var req request.EducationPageListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res := admin.EducationPageListService(req)
	common.ResponseOk(c, res)
}

func (e *Education) Edit(c *gin.Context) {
	var req request.EducationEditRequest
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
	err = admin.EducationEditService(req)
	if err != nil {
		common.ResponseFailWithMsg(c, "添加失败")
		return
	}
	common.ResponseOk(c, nil)
}

func (e *Education) Select(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	res, err := admin.EducationSelectService(id)
	if err != nil {
		common.ResponseFailWithMsg(c, "查询失败")
		return
	}
	common.ResponseOk(c, res)
}

func (e *Education) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		common.ResponseFailWithCode(c, common.ParameterValidError)
		return
	}
	err = admin.EducationDeleteService(id)
	if err != nil {
		common.ResponseFailWithMsg(c, "删除失败")
		return
	}
	common.ResponseOk(c, nil)
}
