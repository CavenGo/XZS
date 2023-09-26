package response

type ExamPaperAnswerPageResponseVM struct {
	Id              int    `json:"id"`
	CreateTime      string `json:"createTime"`
	UserScore       string `json:"userScore"`
	SubjectName     string `json:"subjectName"`
	SubjectId       int    `json:"subjectId"`
	QuestionCount   int    `json:"questionCount"`
	QuestionCorrect int    `json:"questionCorrect"`
	PaperScore      string `json:"paperScore"`
	DoTime          string `json:"doTime"`
	PaperType       int    `json:"paperType"`
	SystemScore     string `json:"systemScore"`
	Status          int    `json:"status"`
	PaperName       string `json:"paperName"`
	UserName        string `json:"userName"`
}

type PageJudgeListResponse struct {
	BasePageResponse
	List []ExamPaperAnswerPageResponseVM `json:"list"`
}
