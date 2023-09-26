package entity

import "time"

type TextContent struct {
	Id         int       `json:"id" gorm:"column:id"`
	Content    string    `json:"content" gorm:"column:content"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`
}
