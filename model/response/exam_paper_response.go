package response

type ExamPaperPage struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	QuestionCount      int    `json:"questionCount"`
	Score              int    `json:"score"`
	CreateTime         string `json:"createTime"`
	CreateUser         int    `json:"createUser"`
	SubjectId          int    `json:"subjectId"`
	PaperType          int    `json:"paperType"`
	FrameTextContentId int    `json:"frameTextContentId"`
}

type ExamPaperPageResponse struct {
	BasePageResponse
	List []ExamPaperPage `json:"list"`
}
