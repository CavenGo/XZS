package model

import (
	"go.uber.org/zap"
	"xzs/global"
	"xzs/model/entity"
)

func ExamPaperAnswerAllCount() int64 {
	var count int64
	global.Db.Model(&entity.ExamPaperAnswer{}).Count(&count)
	return count
}

func ExamPaperAnswerPageList(pageIndex, pageSize, subjectId int) (res []entity.ExamPaperAnswer) {
	offset := (pageIndex - 1) * pageSize
	newDb := global.Db.Model(&entity.ExamPaperAnswer{}).Limit(pageSize).Offset(offset).Order("id desc")
	if subjectId != 0 {
		newDb.Where("subject_id = ?", subjectId)
	}
	err := newDb.Find(&res).Error
	if err != nil {
		zap.L().Info("ExamPaperAnswerPageList err", zap.Error(err))
	}
	return
}

func ExamPaperAnswerPageListAllCount(subjectId int) (count int64) {
	newDb := global.Db.Model(&entity.ExamPaperAnswer{})
	if subjectId != 0 {
		newDb.Where("subject_id = ?", subjectId)
	}
	err := newDb.Count(&count).Error
	if err != nil {
		zap.L().Info("ExamPaperAnswerPageListAllCount err", zap.Error(err))
	}
	return
}
