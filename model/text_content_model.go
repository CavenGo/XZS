package model

import (
	"xzs/global"
	"xzs/model/entity"
)

func TextContentSelectById(id int) (res entity.TextContent, err error) {
	err = global.Db.First(&res, id).Error
	return
}

func TextContentInsertOne(content *entity.TextContent) error {
	return global.Db.Create(content).Error
}

func TextContentUpdateContentById(content string, id int) error {
	return global.Db.Model(&entity.TextContent{}).Where("id = ?", id).Update("content", content).Error
}
