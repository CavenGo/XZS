package request

type ExamPaperPageRequest struct {
	Id         int `json:"id"`
	Level      int `json:"level"`
	SubjectId  int `json:"subjectId"`
	PageIndex  int `json:"pageIndex"`
	PageSize   int `json:"pageSize"`
	PaperType  int `json:"paperType"`
	TaskExamId int `json:"taskExamId"`
}

type ExamPaperEditRequest struct {
	Id            int                    `json:"id"`
	Level         int                    `json:"level"`
	SubjectId     int                    `json:"subjectId"`
	PaperType     int                    `json:"paperType"`
	Name          string                 `json:"name"`
	SuggestTime   int                    `json:"suggestTime"`
	LimitDateTime []string               `json:"limitDateTime"`
	TitleItems    []ExamPaperTitleItemVM `json:"titleItems"`
	Score         string                 `json:"score"`
}

// ExamPaperTitleItemObject 保存到数据库中的题目结构，较为简单
type ExamPaperTitleItemObject struct {
	Name          string                        `json:"name"`
	QuestionItems []ExamPaperQuestionItemObject `json:"questionItems"`
}

// ExamPaperTitleItemVM 返回前端的题目结构，字段较多
type ExamPaperTitleItemVM struct {
	Name          string                `json:"name"`
	QuestionItems []QuestionEditRequest `json:"questionItems"`
}
