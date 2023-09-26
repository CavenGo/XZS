package request

import "xzs/model/response"

type TaskPageListRequest struct {
	GradeLevel int `json:"gradeLevel"`
	PageIndex  int `json:"pageIndex"`
	PageSize   int `json:"pageSize"`
}

type TaskRequestVM struct {
	Id         int                      `json:"id"`
	GradeLevel int                      `json:"gradeLevel"`
	Title      string                   `json:"title"`
	PaperItems []response.ExamPaperPage `json:"paperItems"`
}
