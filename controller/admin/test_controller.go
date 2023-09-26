package admin

import (
	"github.com/gin-gonic/gin"
	"xzs/common"
	"xzs/model"
	"xzs/pkg/dateutil"
)

type Test struct {
}

var TestApi = new(Test)

func (t *Test) Test(c *gin.Context) {
	list := model.SelectEventLogByStartEnd(dateutil.GetMonthStartDay(), dateutil.GetMonthEndDay())
	common.ResponseOk(c, list)
}
