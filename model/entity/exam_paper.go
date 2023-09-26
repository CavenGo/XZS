package entity

import "time"

type ExamPaper struct {
	Id                 int       `json:"id" gorm:"column:id"`
	Name               string    `json:"name" gorm:"column:name"`
	SubjectId          int       `json:"subjectId" gorm:"column:subject_id"`
	PaperType          int       `json:"paperType" gorm:"column:paper_type"`
	GradeLevel         int       `json:"gradeLevel" gorm:"column:grade_level"`
	Score              int       `json:"score" gorm:"column:score"`
	QuestionCount      int       `json:"questionCount" gorm:"column:question_count"`
	SuggestTime        int       `json:"suggestTime" gorm:"column:suggest_time"`
	LimitStartTime     time.Time `json:"limitStartTime" gorm:"column:limit_start_time"`
	LimitEndTime       time.Time `json:"limitEndTime" gorm:"column:limit_end_time"`
	FrameTextContentId int       `json:"frameTextContentId" gorm:"column:frame_text_content_id"`
	CreateUser         int       `json:"createUser" gorm:"column:create_user"`
	CreateTime         time.Time `json:"createTime" gorm:"column:create_time"`
	Deleted            bool      `json:"deleted" gorm:"column:deleted"`
	TaskExamId         int       `json:"taskExamId" gorm:"column:task_exam_id"`
}
