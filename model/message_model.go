package model

import (
	"go.uber.org/zap"
	"xzs/global"
	"xzs/model/entity"
)

func MessageInsert(model *entity.Message) error {
	err := global.Db.Model(&entity.Message{}).Create(&model).Error
	return err
}

func MessagePaperPageList(pageIndex, pageSize int, sendUserName string) (res []entity.Message) {
	offset := (pageIndex - 1) * pageSize
	newDb := global.Db.Model(&entity.Message{}).Limit(pageSize).Offset(offset).Order("id desc")

	if sendUserName != "" {
		newDb.Where("send_user_name like ?", "%"+sendUserName+"%")
	}
	err := newDb.Find(&res).Error
	if err != nil {
		zap.L().Info("MessagePaperPageList err", zap.Error(err))
	}
	return
}

func MessagePageListAllCount(sendUserName string) (count int64) {
	newDb := global.Db.Model(&entity.Message{})
	if sendUserName != "" {
		newDb.Where("send_user_name like ?", "%"+sendUserName+"%")
	}
	err := newDb.Count(&count).Error
	if err != nil {
		zap.L().Info("MessagePageListAllCount err", zap.Error(err))
	}
	return
}
