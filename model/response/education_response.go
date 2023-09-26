package response

type PageListSubject struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Level     int    `json:"level"`
	LevelName string `json:"levelName"`
}

type EducationPageListResponse struct {
	BasePageResponse
	List []PageListSubject `json:"list"`
}
