package model

import (
	"go.uber.org/zap"
	"xzs/global"
	"xzs/model/entity"
)

func AddUserEventLog(userEventLog *entity.UserEventLog) (err error) {
	err = global.Db.Create(userEventLog).Error
	return
}

func SelectEventLogByStartEnd(start, end string) []entity.UserEventLogCount {
	var res []entity.UserEventLogCount
	err := global.Db.Raw("SELECT create_time as `name`,COUNT(create_time) as `value` from (SELECT DATE_FORMAT(create_time,'%Y-%m-%d') as create_time from t_user_event_log WHERE  create_time  between  ?  and  ?) a GROUP BY create_time", start, end).Scan(&res).Error
	if err != nil {
		zap.L().Error("SelectEventLogByStartEnd", zap.Error(err))
	}
	return res
}

func UserEventLogPageList(pageIndex, pageSize, userId int, userName string) (res []entity.UserEventLog) {
	offset := (pageIndex - 1) * pageSize
	newDb := global.Db.Model(&entity.UserEventLog{}).Limit(pageSize).Offset(offset).Order("id desc")
	if userId != 0 {
		newDb.Where("user_id = ?", userId)
	}
	if userName != "" {
		newDb.Where("user_name = ?", userName)
	}
	err := newDb.Find(&res).Error
	if err != nil {
		zap.L().Info("UserEventLOgPageList err", zap.Error(err))
	}
	return
}

func UserEventLogPageListAllCount(userId int, userName string) (count int64) {
	newDb := global.Db.Model(&entity.UserEventLog{})
	if userId != 0 {
		newDb.Where("user_id = ?", userId)
	}
	if userName != "" {
		newDb.Where("user_name = ?", userName)
	}
	err := newDb.Count(&count).Error
	if err != nil {
		zap.L().Info("UserEventLogPageListAllCount err", zap.Error(err))
	}
	return
}
