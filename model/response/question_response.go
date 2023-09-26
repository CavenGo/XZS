package response

type QuestionResponseVM struct {
	Id                   int         `json:"id"`
	QuestionType         int         `json:"questionType"`
	TextContentId        interface{} `json:"textContentId"`
	CreateTime           string      `json:"createTime"`
	SubjectId            int         `json:"subjectId"`
	CreateUser           int         `json:"createUser"`
	Score                string      `json:"score"`
	Status               int         `json:"status"`
	Correct              string      `json:"correct"`
	AnalyzeTextContentId interface{} `json:"analyzeTextContentId"`
	Difficult            int         `json:"difficult"`
	ShortTitle           string      `json:"shortTitle"`
}

type QuestionPageListResponse struct {
	BasePageResponse
	List []QuestionResponseVM `json:"list"`
}
