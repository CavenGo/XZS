package model

import (
	"go.uber.org/zap"
	"xzs/global"
	"xzs/model/entity"
)

func TaskExamPageList(pageIndex, pageSize, gradeLevel int) (res []entity.TaskExam) {
	offset := (pageIndex - 1) * pageSize
	newDb := global.Db.Model(&entity.TaskExam{}).Where("deleted = 0").Limit(pageSize).Offset(offset).Order("id desc")
	if gradeLevel != 0 {
		newDb.Where("grade_level = ?", gradeLevel)
	}
	err := newDb.Find(&res).Error
	if err != nil {
		zap.L().Info("TaskExamPageList err", zap.Error(err))
	}
	return
}

func TaskExamPageListAllCount(gradeLevel int) (count int64) {
	newDb := global.Db.Model(&entity.ExamPaper{}).Where("deleted = 0")
	if gradeLevel != 0 {
		newDb.Where("grade_level = ?", gradeLevel)
	}
	err := newDb.Count(&count).Error
	if err != nil {
		zap.L().Info("TaskExamPageListAllCount err", zap.Error(err))
	}
	return
}

func TaskExamDeleteById(id int) error {
	return global.Db.Model(&entity.TaskExam{}).Where("id = ?", id).Update("deleted", 1).Error
}

func TaskSelectById(id int) (res entity.TaskExam) {
	global.Db.First(&res, id)
	return
}

func TaskExamUpdateById(id int, model entity.TaskExam) error {
	return global.Db.Model(&entity.TaskExam{}).Where("id = ?", id).Updates(model).Error
}

func TaskExamInsertOne(model *entity.TaskExam) error {
	return global.Db.Create(model).Error
}
