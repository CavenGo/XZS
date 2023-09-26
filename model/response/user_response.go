package response

type PageListUser struct {
	Id             int    `json:"id"`
	UserUuid       string `json:"userUuid"`
	UserName       string `json:"userName"`
	RealName       string `json:"realName"`
	Age            int    `json:"age"`
	Role           int    `json:"role"`
	Sex            int    `json:"sex"`
	BirthDay       string `json:"birthDay"`
	Phone          string `json:"phone"`
	LastActiveTime string `json:"lastActiveTime"`
	CreateTime     string `json:"createTime"`
	ModifyTime     string `json:"modifyTime"`
	Status         int    `json:"status"`
	UserLevel      int    `json:"userLevel"`
	ImagePath      string `json:"imagePath"`
}

type PageListResponse struct {
	BasePageResponse
	List []PageListUser `json:"list"`
}

type UserEventLogVM struct {
	Id         int    `json:"id"`
	UserId     int    `json:"userId"`
	UserName   string `json:"userName"`
	RealName   string `json:"realName"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
}

type UserEventLogResponse struct {
	BasePageResponse
	List []UserEventLogVM `json:"list"`
}
