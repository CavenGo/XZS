package entity

import "time"

type ExamPaperQuestionCustomerAnswer struct {
	Id                    int       `json:"id" gorm:"column:id"`
	QuestionId            int       `json:"questionId" gorm:"column:question_id"`
	ExamPaperId           int       `json:"examPaperId" gorm:"column:exam_paper_id"`
	ExamPaperAnswerId     int       `json:"examPaperAnswerId" gorm:"column:exam_paper_answer_id"`
	QuestionType          int       `json:"questionType" gorm:"column:question_type"`
	SubjectId             int       `json:"subjectId" gorm:"column:subject_id"`
	CustomerScore         int       `json:"customerScore" gorm:"column:customer_score"`
	QuestionScore         int       `json:"questionScore" gorm:"column:question_score"`
	QuestionTextContentId int       `json:"questionTextContentId" gorm:"column:question_text_content_id"`
	Answer                string    `json:"answer" gorm:"column:answer"`
	TextContentId         int       `json:"textContentId" gorm:"column:text_content_id"`
	DoRight               byte      `json:"doRight" gorm:"column:do_right"`
	CreateUser            int       `json:"createUser" gorm:"column:create_user"`
	CreateTime            time.Time `json:"createTime" gorm:"column:create_time"`
	ItemOrder             int       `json:"itemOrder" gorm:"column:item_order"`
}

type ExamPaperQuestionCustomerAnswerCount struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
