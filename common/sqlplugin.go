package common

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xzs/config"
)

const callBackAfterName = "core:after"

type SqlTracePlugin struct{}

func (s *SqlTracePlugin) Name() string {
	return "sqlTracePlugin"
}

func (s *SqlTracePlugin) Initialize(db *gorm.DB) (err error) {
	db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

func after(db *gorm.DB) {
	if !config.GlobalConf.Mysql.SqlLog {
		return
	}
	dbCtx := db.Statement.Context
	ctx, ok := dbCtx.(*gin.Context)
	if !ok {
		return
	}

	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	ctx.Set("sql", sql)
}
