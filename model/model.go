package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"xzs/common"
	"xzs/config"
)

func InitMysql() (db *gorm.DB, err error) {
		/*
		config.GlobalConf.Mysql.User = "root"
		config.GlobalConf.Mysql.Password = "111111"
		config.GlobalConf.Mysql.Host = "127.0.0.1"
		config.GlobalConf.Mysql.Port = 3306
		config.GlobalConf.Mysql.Dbname = "artist"
		*/
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GlobalConf.Mysql.User,
		config.GlobalConf.Mysql.Password,
		config.GlobalConf.Mysql.Host,
		config.GlobalConf.Mysql.Port,
		config.GlobalConf.Mysql.Dbname,
	)
	//dsn := fmt.Sprintf("root:111111@tcp(127.0.0.1:3306)/artist?charset=utf8mb4&parseTime=True&loc=Local",)
	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 指名表前缀
			SingularTable: true, // 使用单数表名
		},
	}
	if config.GlobalConf.Mysql.LogMode {
		// 开启调试模式，打印所有的sql语句
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err = gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		return
	}
	sqlDb.SetMaxIdleConns(config.GlobalConf.Mysql.MaxIdleConns)
	sqlDb.SetMaxOpenConns(config.GlobalConf.Mysql.MaxOpenConns)

	err = db.Use(&common.SqlTracePlugin{})
	return
}
