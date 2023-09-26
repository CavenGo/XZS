package admin

import (
	"github.com/gin-gonic/gin"
	"xzs/common"
	"xzs/service/admin"
)

type Dashboard struct {
}

var DashoboardApi = new(Dashboard)

// Index 后台首页接口
func (d *Dashboard) Index(c *gin.Context) {
	res := admin.DashboardIndexService(c)
	common.ResponseOk(c, res)
}
