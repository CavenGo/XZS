package entity

import "time"

type Question struct {
	Id                int       `json:"id" gorm:"column:id"`
	QuestionType      int       `json:"questionType" gorm:"column:question_type"`
	SubjectId         int       `json:"subjectId" gorm:"column:subject_id"`
	Score             int       `json:"score" gorm:"column:score"`
	GradeLevel        int       `json:"gradeLevel" gorm:"column:grade_level"`
	Difficult         int       `json:"difficult" gorm:"column:difficult"`
	Correct           string    `json:"correct" gorm:"column:correct"`
	InfoTextContentId int       `json:"infoTextContentId" gorm:"column:info_text_content_id"`
	CreateUser        int       `json:"createUser" gorm:"column:create_user"`
	Status            int       `json:"status" gorm:"column:status"`
	CreateTime        time.Time `json:"createTime" gorm:"column:create_time"`
	Deleted           bool      `json:"deleted" gorm:"column:deleted"`
}
