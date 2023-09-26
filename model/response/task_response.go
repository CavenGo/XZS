package response

type TaskPageResponseVM struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	GradeLevel     int    `json:"gradeLevel"`
	CreateUserName string `json:"createUserName"`
	CreateTime     string `json:"createTime"`
	Deleted        bool   `json:"deleted"`
}

type TaskPageListResponse struct {
	BasePageResponse
	List []TaskPageResponseVM `json:"list"`
}

type TaskRequestVM struct {
	Id         int             `json:"id"`
	GradeLevel int             `json:"gradeLevel"`
	Title      string          `json:"title"`
	PaperItems []ExamPaperPage `json:"paperItems" validate:"required,gte=1"`
}

type TaskItemObject struct {
	ExamPaperId   int    `json:"examPaperId"`
	ExamPaperName string `json:"examPaperName"`
	ItemOrder     int    `json:"itemOrder"`
}
