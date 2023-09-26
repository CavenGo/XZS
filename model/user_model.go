package model

import (
	"go.uber.org/zap"
	"xzs/global"
	"xzs/model/entity"
)

// FindUserByUserName 根据用户名查询用户
func FindUserByUserName(userName string) (user entity.User, err error) {
	err = global.Db.Where("deleted = 0 and user_name = ?", userName).First(&user).Error
	return
}

// GetUserById 根据ID获取用户信息
func GetUserById(id int) (user entity.User, err error) {
	err = global.Db.Where("id = ?", id).First(&user).Error
	return
}

// AddUser 添加用户
func AddUser(user *entity.User) (err error) {
	err = global.Db.Create(user).Error
	return err
}

func UserPageList(pageIndex, pageSize, role int, userName string) []entity.User {
	var res []entity.User
	offset := (pageIndex - 1) * pageSize
	newDb := global.Db.Limit(pageSize).Offset(offset).Where("role = ?", role)
	if userName != "" {
		newDb.Where("user_name like ?", "%"+userName+"%")
	}
	err := newDb.Order("id desc").Find(&res).Error
	if err != nil {
		zap.L().Info("UserPageList，", zap.Any("pageIndex", pageSize), zap.Any("pageSize", pageSize), zap.Any("userName", userName), zap.Error(err))
	}
	return res
}

func UserAllCountByUserNameRole(userName string, role int) (count int64) {
	newDb := global.Db.Model(&entity.User{}).Where("role = ?", role).Count(&count)
	if userName != "" {
		newDb.Where("user_name like ?", "%"+userName+"%")
	}
	return
}

func UpdateStatusById(id, status int) (err error) {
	err = global.Db.Model(&entity.User{}).Where("id = ?", id).Update("status", status).Error
	return err
}

func UpdateUserById(id int, user entity.User) (err error) {
	err = global.Db.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error
	return
}

func UpdateUserByUserName(userName string, user entity.User) error {
	err := global.Db.Model(&entity.User{}).Where("user_name = ?", userName).Updates(user).Error
	return err
}

func DeleteUserById(id int) error {
	err := global.Db.Model(&entity.User{}).Where("id = ?", id).Updates(entity.User{
		Deleted: true,
	}).Error
	return err
}

func SelectUserKeyValueByUserName(userName string) (list []entity.UserKeyValue) {
	err := global.Db.Model(&entity.User{}).Select("id as value, user_name as name").Where("deleted = 0 and user_name like ?", "%"+userName+"%").Limit(5).Find(&list).Error
	if err != nil {
		zap.L().Info("SelectUserKeyValueByUserName err", zap.String("userName", userName), zap.Error(err))
	}
	return
}

// FindUserByIds 根据id批量查询
func FindUserByIds(ids []int) (list []entity.User) {
	global.Db.Find(&list, ids)
	return
}
