package entity

import "time"

type MessageUser struct {
	ID              int        `json:"id" gorm:"column:id"`
	MessageID       int        `json:"messageId" gorm:"column:message_id"`
	ReceiveUserID   int        `json:"receiveUserId" gorm:"column:receive_user_id"`
	ReceiveUserName string     `json:"receiveUserName" gorm:"column:receive_user_name"`
	ReceiveRealName string     `json:"receiveRealName" gorm:"column:receive_real_name"`
	Readed          uint8      `json:"readed" gorm:"column:readed"`
	CreateTime      time.Time  `json:"createTime" gorm:"column:create_time"`
	ReadTime        *time.Time `json:"readTime" gorm:"column:read_time"`
}
