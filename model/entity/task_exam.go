package entity

import "time"

type TaskExam struct {
	Id                 int       `json:"id" gorm:"column:id"`
	Title              string    `json:"title" gorm:"column:title"`
	GradeLevel         int       `json:"gradeLevel" gorm:"column:grade_level"`
	FrameTextContentId int       `json:"frameTextContentId" gorm:"column:frame_text_content_id"`
	CreateUser         int       `json:"createUser" gorm:"column:create_user"`
	CreateTime         time.Time `json:"createTime" gorm:"column:create_time"`
	Deleted            bool      `json:"deleted" gorm:"column:deleted"`
	CreateUserName     string    `json:"createUserName" gorm:"column:create_user_name"`
}
