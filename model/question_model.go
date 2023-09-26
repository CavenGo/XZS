package model

import (
	"go.uber.org/zap"
	"xzs/global"
	"xzs/model/entity"
)

func QuestionSelectByIds(ids []int) (list []entity.Question) {
	if len(ids) == 0 {
		return
	}
	err := global.Db.Find(&list, ids).Error
	if err != nil {
		zap.L().Info("QuestionSelectByIds err", zap.Ints("ids", ids), zap.Error(err))
	}
	return
}

func QuestionAllCount() int64 {
	var count int64
	global.Db.Model(&entity.Question{}).Where("deleted = ?", 0).Count(&count)
	return count
}

func QuestionDeleteById(id int) error {
	return global.Db.Model(&entity.Question{}).Where("id = ?", id).Update("deleted", 1).Error
}

func QuestionSelectById(id int) (res entity.Question) {
	global.Db.First(&res, id)
	return
}

func QuestionPageList(pageIndex, pageSize, id, level, subjectId, questionType int) (res []entity.Question) {
	offset := (pageIndex - 1) * pageSize
	newDb := global.Db.Model(&entity.Question{}).Where("deleted = 0").Limit(pageSize).Offset(offset).Order("id desc")
	if id != 0 {
		newDb.Where("id = ?", id)
	}
	if level != 0 {
		newDb.Where("grade_level = ?", level)
	}
	if subjectId != 0 {
		newDb.Where("subject_id = ?", subjectId)
	}
	if questionType != 0 {
		newDb.Where("question_type = ?", questionType)
	}
	err := newDb.Find(&res).Error
	if err != nil {
		zap.L().Info("QuestionPageList err", zap.Error(err))
	}
	return
}

func QuestionPageListAllCount(id, level, subjectId, questionType int) (count int64) {
	newDb := global.Db.Model(&entity.Question{}).Where("deleted = 0")
	if id != 0 {
		newDb.Where("id = ?", id)
	}
	if level != 0 {
		newDb.Where("grade_level = ?", level)
	}
	if subjectId != 0 {
		newDb.Where("subject_id = ?", subjectId)
	}
	if questionType != 0 {
		newDb.Where("question_type = ?", questionType)
	}
	err := newDb.Count(&count).Error
	if err != nil {
		zap.L().Info("QuestionPageListAllCount err", zap.Error(err))
	}
	return
}

func QuestionInsertOne(model *entity.Question) error {
	return global.Db.Create(model).Error
}

func QuestionUpdateById(id int, model entity.Question) error {
	return global.Db.Model(&entity.Question{}).Where("id = ?", id).Updates(model).Error
}
