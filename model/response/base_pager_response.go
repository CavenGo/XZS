package response

type BasePageResponse struct {
	Total             int   `json:"total"`
	PageNum           int   `json:"pageNum"`
	PageSize          int   `json:"pageSize"`
	Size              int   `json:"size"`
	StartRow          int   `json:"startRow"`
	EndRow            int   `json:"endRow"`
	Pages             int   `json:"pages"`
	PrePage           int   `json:"prePage"`
	NextPage          int   `json:"nextPage"`
	IsFirstPage       bool  `json:"isFirstPage"`
	IsLastPage        bool  `json:"isLastPage"`
	HasPreviousPage   bool  `json:"hasPreviousPage"`
	HasNextPage       bool  `json:"hasNextPage"`
	NavigatePages     int   `json:"navigatePages"`
	NavigatepageNums  []int `json:"navigatepageNums"`
	NavigateFirstPage int   `json:"navigateFirstPage"`
	NavigateLastPage  int   `json:"navigateLastPage"`
}
