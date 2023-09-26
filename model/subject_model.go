package model

import (
	"go.uber.org/zap"
	"xzs/config"
	"xzs/global"
	"xzs/model/entity"
)

func SelectAllSubject() (list []entity.Subject) {
	err := global.Db.Debug().Find(&list).Error
	sql := global.Db.Dialector.Explain(global.Db.Statement.SQL.String(), global.Db.Statement.Vars...)
	config.ZapLogger.Info("SelectAllSubject", zap.String("sql", sql))
	if err != nil {
		zap.L().Info("SelectAllSubject error", zap.Error(err))
	}
	return
}

func SubjectPageList(pageIndex, pageSize, id, level int) (res []entity.Subject) {
	offset := (pageIndex - 1) * pageSize
	newDb := global.Db.Model(&entity.Subject{}).Where("deleted = 0").Limit(pageSize).Offset(offset).Order("id desc")
	if id != 0 {
		newDb.Where("id = ?", id)
	}
	if level != 0 {
		newDb.Where("level = ?", level)
	}
	err := newDb.Find(&res).Error
	if err != nil {
		zap.L().Info("SubjectPageList err", zap.Any("pageIndex", pageSize), zap.Any("pageSize", pageSize), zap.Int("id", id), zap.Int("level", level), zap.Error(err))
	}
	return

}

func SubjectAllCountByUserNameRole(id, level int) (count int64) {
	newDb := global.Db.Model(&entity.Subject{}).Where("deleted = 0")
	if id != 0 {
		newDb.Where("id = ?", id)
	}
	if level != 0 {
		newDb.Where("level = ?", level)
	}
	newDb.Count(&count)
	return
}

func AddSubject(model *entity.Subject) error {
	err := global.Db.Create(model).Error
	return err
}

func UpdateSubjectById(id int, model entity.Subject) error {
	err := global.Db.Model(&entity.Subject{}).Where("id = ?", id).Updates(model).Error
	return err
}

func SelectSubjectById(id int) (subject entity.Subject, err error) {
	err = global.Db.Where("id = ?", id).First(&subject).Error
	return
}

func DeleteSubjectById(id int) error {
	err := global.Db.Where("id = ?", id).Update("deleted", 1).Error
	return err
}
