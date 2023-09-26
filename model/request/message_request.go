package request

type MessageSendVM struct {
	Title          string `json:"title"`
	Content        string `json:"content"`
	ReceiveUserIds []int  `json:"receiveUserIds" validate:"required,min=1"`
}

type MessagePageRequestVM struct {
	SendUserName string `json:"sendUserName"`
	PageIndex    int    `json:"pageIndex"`
	PageSize     int    `json:"pageSize"`
}
