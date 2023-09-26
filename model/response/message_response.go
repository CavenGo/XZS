package response

type MessageResponseVM struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	SendUserName     string `json:"sendUserName"`
	Receives         string `json:"receives"`
	ReceiveUserCount int    `json:"receiveUserCount"`
	ReadCount        int    `json:"readCount"`
	CreateTime       string `json:"createTime"`
}

type MessagePageListResponse struct {
	BasePageResponse
	List []MessageResponseVM `json:"list"`
}
