package request

type QuestionEditVm struct {
	Id           int                  `json:"id"`
	QuestionType int                  `json:"questionType"`
	SubjectId    int                  `json:"subjectId"`
	Title        string               `json:"title"`
	GradeLevel   int                  `json:"gradeLevel"`
	Items        []QuestionEditItemVm `json:"items"`
}

type QuestionEditItemVm struct {
	Prefix   string `json:"prefix"`
	Content  string `json:"content"`
	Score    int    `json:"score"`
	ItemUuid string `json:"itemUuid"`
}

type QuestionEditRequest struct {
	Id           int                  `json:"id"`
	QuestionType int                  `json:"questionType"`
	SubjectId    int                  `json:"subjectId"`
	Title        string               `json:"title" validate:"required,gte=1,lte=5"`
	GradeLevel   int                  `json:"gradeLevel"`
	Items        []QuestionEditItemVm `json:"items"`
	Analyze      string               `json:"analyze"`
	CorrectArray []string             `json:"correctArray"`
	Correct      string               `json:"correct"`
	Score        string               `json:"score"`
	Difficult    int                  `json:"difficult" validate:"required"`
	ItemOrder    int                  `json:"itemOrder"`
}

// QuestionObject 题目结构解析结构体
type QuestionObject struct {
	TitleContent        string               `json:"titleContent"`
	Analyze             string               `json:"analyze"`
	Correct             string               `json:"correct"`
	QuestionItemObjects []QuestionEditItemVm `json:"questionItemObjects"`
}

type ExamPaperQuestionItemObject struct {
	Id        int `json:"id"`
	ItemOrder int `json:"itemOrder"`
}

type QuestionPageRequestVM struct {
	Id           int `json:"id"`
	QuestionType int `json:"questionType"`
	Level        int `json:"level"`
	SubjectId    int `json:"subjectId"`
	PageIndex    int `json:"pageIndex"`
	PageSize     int `json:"pageSize"`
}
