package request

type EducationPageListRequest struct {
	Id        int `json:"id"`
	Level     int `json:"level"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type EducationEditRequest struct {
	Id        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	Level     int    `json:"level" validate:"required"`
	LevelName string `json:"levelName" validate:"required"`
}
