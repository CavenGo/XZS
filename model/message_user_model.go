package model

import (
	"xzs/global"
	"xzs/model/entity"
)

func MessageUserBatchSave(list *[]entity.MessageUser) (err error) {
	err = global.Db.Create(list).Error
	return err
}

func MessageUserFindList(where map[string]interface{}) (list []entity.MessageUser) {
	global.Db.Model(&entity.MessageUser{}).Where(where).Find(&list)
	return
}
