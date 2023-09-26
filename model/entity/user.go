package entity

import "time"

type User struct {
	Id             int       `json:"id" gorm:"column:id"`
	UserUuid       string    `json:"userUUid" gorm:"column:user_uuid"`
	UserName       string    `json:"userName" gorm:"column:user_name"`
	Password       string    `json:"password" gorm:"column:password"`
	RealName       string    `json:"realName" gorm:"column:real_name"`
	Age            int       `json:"age" gorm:"column:age"`
	Sex            int       `json:"sex" gorm:"column:sex"`
	BirthDay       time.Time `json:"BirthDay" gorm:"column:birth_day"`
	UserLevel      int       `json:"userLevel" gorm:"column:user_level"`
	Phone          string    `json:"phone" gorm:"column:phone"`
	Role           int       `json:"role" gorm:"column:role"`
	Status         int       `json:"status" gorm:"column:status"`
	ImagePath      string    `json:"ImagePath" gorm:"column:image_path"`
	CreateTime     time.Time `json:"CreateTime" gorm:"column:create_time"`
	ModifyTime     time.Time `json:"ModifyTime" gorm:"column:modify_time"`
	LastActiveTime time.Time `json:"LastActiveTime" gorm:"column:last_active_time"`
	Deleted        bool      `json:"deleted" gorm:"column:deleted"`
	WxOpenId       string    `json:"WxOpenId" gorm:"column:wx_open_id"`
}

type UserKeyValue struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
