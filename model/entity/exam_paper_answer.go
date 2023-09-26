package entity

import "time"

type ExamPaperAnswer struct {
	Id              int       `json:"id" gorm:"column:id"`
	ExamPaperId     int       `json:"examPaperId" gorm:"column:exam_paper_id"`
	PaperName       string    `json:"paperName" gorm:"column:paper_name"`
	PaperType       int       `json:"paperType" gorm:"column:paper_type"`
	SubjectId       int       `json:"subjectId" gorm:"column:subject_id"`
	SystemScore     int       `json:"systemScore" gorm:"column:system_score"`
	UserScore       int       `json:"userScore" gorm:"column:user_score"`
	PaperScore      int       `json:"paperScore" gorm:"column:paper_score"`
	QuestionCorrect int       `json:"questionCorrect" gorm:"column:question_correct"`
	QuestionCount   int       `json:"questionCount" gorm:"column:question_count"`
	DoTime          int       `json:"doTime" gorm:"column:do_time"`
	Status          int       `json:"status" gorm:"column:status"`
	CreateUser      int       `json:"createUser" gorm:"column:create_user"`
	CreateTime      time.Time `json:"createTime" gorm:"column:create_time"`
	TaskExamId      int       `json:"taskExamId" gorm:"column:task_exam_id"`
}
