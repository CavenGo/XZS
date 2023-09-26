package model

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"xzs/global"
	"xzs/model/entity"
)

func ExamPaperAllCount(c *gin.Context) int64 {
	var count int64
	global.Db.WithContext(c).Model(&entity.ExamPaper{}).Where("deleted = ?", 0).Count(&count)
	return count
}

func ExamPaperPageList(pageIndex, pageSize, id, level, subjectId, paperType int) (res []entity.ExamPaper) {
	offset := (pageIndex - 1) * pageSize
	newDb := global.Db.Model(&entity.ExamPaper{}).Where("deleted = 0").Limit(pageSize).Offset(offset).Order("id desc")
	if id != 0 {
		newDb.Where("id = ?", id)
	}
	if level != 0 {
		newDb.Where("grade_level = ?", level)
	}
	if subjectId != 0 {
		newDb.Where("subject_id = ?", subjectId)
	}
	if paperType != 0 {
		newDb.Where("paper_type = ?", paperType)
	}
	err := newDb.Find(&res).Error
	if err != nil {
		zap.L().Info("ExamPaperPageList err", zap.Error(err))
	}
	return
}

func ExamPaperPageListAllCount(id, level, subjectId, paperType int) (count int64) {
	newDb := global.Db.Model(&entity.ExamPaper{}).Where("deleted = 0 and task_exam_id = 0")
	if id != 0 {
		newDb.Where("id = ?", id)
	}
	if level != 0 {
		newDb.Where("grade_level = ?", level)
	}
	if subjectId != 0 {
		newDb.Where("subject_id = ?", subjectId)
	}
	if paperType != 0 {
		newDb.Where("paper_type = ?", paperType)
	}
	err := newDb.Count(&count).Error
	if err != nil {
		zap.L().Info("ExamPaperPageListAllCount err", zap.Error(err))
	}
	return
}

func TaskExamPaperPageList(pageIndex, pageSize, level, paperType int) (res []entity.ExamPaper) {
	offset := (pageIndex - 1) * pageSize
	newDb := global.Db.Model(&entity.ExamPaper{}).Where("deleted = 0 and task_exam_id = 0").Limit(pageSize).Offset(offset).Order("id desc")

	if level != 0 {
		newDb.Where("grade_level = ?", level)
	}
	if paperType != 0 {
		newDb.Where("paper_type = ?", paperType)
	}
	err := newDb.Find(&res).Error
	if err != nil {
		zap.L().Info("ExamPaperPageList err", zap.Error(err))
	}
	return
}

func TaskExamPaperPageListAllCount(level, paperType int) (count int64) {
	newDb := global.Db.Model(&entity.ExamPaper{}).Where("deleted = 0 and task_exam_id is null")
	if level != 0 {
		newDb.Where("grade_level = ?", level)
	}
	if paperType != 0 {
		newDb.Where("paper_type = ?", paperType)
	}
	err := newDb.Count(&count).Error
	if err != nil {
		zap.L().Info("ExamPaperPageListAllCount err", zap.Error(err))
	}
	return
}

func DeleteExamPaperById(id int) error {
	return global.Db.Model(&entity.ExamPaper{}).Where("id = ?", id).Update("deleted", 1).Error
}

func ExamPaperSelectById(id int) (res entity.ExamPaper) {
	err := global.Db.Where("id = ?", id).First(&res).Error
	if err != nil {
		zap.L().Info("ExamPaperSelectById", zap.Int("id", id), zap.Error(err))
	}
	return
}

func ExamPaperUpdateById(examPaper entity.ExamPaper, id int) error {
	return global.Db.Model(&entity.ExamPaper{}).Where("id = ?", id).Updates(examPaper).Error
}

func ExamPaperInsertOne(model *entity.ExamPaper) error {
	return global.Db.Create(model).Error
}

func ExamPaperUpdateTaskId(taskExamId, id int) error {
	return global.Db.Model(&entity.ExamPaper{}).Where("id = ?", id).Update("task_exam_id", taskExamId).Error
}

func ClearTaskExamId(taskExamId int) error {
	return global.Db.Model(&entity.ExamPaper{}).Where("task_exam_id = ?", taskExamId).Update("task_exam_id", 0).Error
}
