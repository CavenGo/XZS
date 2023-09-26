package response

type DashBoardIndexResponse struct {
	ExamPaperCount             int64    `json:"examPaperCount"`
	QuestionCount              int64    `json:"questionCount"`
	DoExamPaperCount           int64    `json:"doExamPaperCount"`
	DoQuestionCount            int64    `json:"doQuestionCount"`
	MothDayUserActionValue     []int    `json:"mothDayUserActionValue"`
	MothDayDoExamQuestionValue []int    `json:"mothDayDoExamQuestionValue"`
	MothDayText                []string `json:"mothDayText"`
}
