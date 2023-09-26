package entity

import "time"

type Message struct {
	ID               int       `json:"id" gorm:"column:id"`
	Title            string    `json:"title" gorm:"column:title"`
	Content          string    `json:"content" gorm:"column:content"`
	CreateTime       time.Time `json:"createTime" gorm:"column:create_time"`
	SendUserID       int       `json:"sendUserId" gorm:"column:send_user_id"`
	SendUserName     string    `json:"sendUserName" gorm:"column:send_user_name"`
	SendRealName     string    `json:"sendRealName" gorm:"column:send_real_name"`
	ReceiveUserCount int       `json:"receiveUserCount" gorm:"column:receive_user_count"`
	ReadCount        int       `json:"readCount" gorm:"column:read_count"`
}
