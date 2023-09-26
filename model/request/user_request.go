package request

type LoginRequest struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	Remember bool   `json:"remember"`
}

type PageListRequest struct {
	UserName  string `json:"userName"`
	Role      int    `json:"role" validate:"required"`
	PageIndex int    `json:"pageIndex" validate:"required"`
	PageSize  int    `json:"pageSize" validate:"required"`
}

type UserEditRequest struct {
	Id        int    `json:"id"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	RealName  string `json:"realName"`
	Age       int    `json:"age"`
	Status    int    `json:"status"`
	Sex       int    `json:"sex"`
	BirthDay  string `json:"birthDay"`
	Phone     string `json:"phone"`
	Role      int    `json:"role"`
	UserLevel int    `json:"userLevel"`
}

type UserUpdateRequest struct {
	RealName string `json:"realName"`
	Phone    string `json:"phone"`
}

type UserNameRequest struct {
	UserName string `json:"userName" validate:"required"`
}

type UserEventPageRequestVM struct {
	UserId    int    `json:"userId"`
	UserName  string `json:"userName"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}
