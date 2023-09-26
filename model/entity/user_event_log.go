package entity

import "time"

type UserEventLog struct {
	Id         int       `json:"id" gorm:"column:id"`
	UserId     int       `json:"userId" gorm:"column:user_id"`
	UserName   string    `json:"userName" gorm:"column:user_name"`
	RealName   string    `json:"RealName" gorm:"column:real_name"`
	Content    string    `json:"content" gorm:"column:content"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`
}

type UserEventLogCount struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
