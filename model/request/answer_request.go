package request

type ExamPaperAnswerPageRequestVM struct {
	SubjectId int `json:"subjectId"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}
