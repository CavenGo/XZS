package admin

import "xzs/model/response"

func CalBasePage(pageIndex, pageSize, count, nowRows int) response.BasePageResponse {
	res := response.BasePageResponse{
		NavigatePages: 8,
		PageNum:       pageIndex,
		PageSize:      pageSize,
		Size:          pageSize,
		Total:         count,
	}
	pages := 0
	// 计算总页数
	if pageSize > 0 {
		pages = count / pageSize
		if count%pageSize != 0 {
			pages++
		}
	}
	if pages < res.NavigatePages {
		res.NavigatePages = pages
	}
	if pageIndex == 1 {
		res.IsFirstPage = true
	} else {
		res.HasPreviousPage = true
		res.PrePage = pageIndex - 1
	}
	if pageIndex == pages {
		res.IsLastPage = true
	}
	if pageIndex < pages {
		res.HasNextPage = true
		res.NextPage = pageIndex + 1
	}
	res.StartRow = pageIndex*pageSize + 1
	res.EndRow = res.StartRow + nowRows

	// 计算导航
	var navigatepageNums []int
	if pages < res.NavigatePages {
		for i := 0; i < pages; i++ {
			navigatepageNums = append(navigatepageNums, i+1)
		}
	} else {
		startNum := res.PageNum - res.NavigatePages/2
		endNum := res.PageNum + res.NavigatePages/2
		if startNum < 1 {
			startNum = 1
			//(最前navigatePages页
			for i := 0; i < res.NavigatePages; i++ {
				navigatepageNums = append(navigatepageNums, startNum)
				startNum++
			}
		} else if endNum > pages {
			endNum = pages
			//最后navigatePages页
			for i := res.NavigatePages - 1; i >= 0; i-- {
				navigatepageNums = append(navigatepageNums, endNum-1)
			}
		} else {
			//所有中间页
			for i := 0; i < res.NavigatePages; i++ {
				navigatepageNums = append(navigatepageNums, startNum)
				startNum++
			}
		}
	}
	res.NavigatepageNums = navigatepageNums

	// 计算前后页
	if len(res.NavigatepageNums) > 0 {
		res.NavigateFirstPage = navigatepageNums[0]
		res.NavigateLastPage = navigatepageNums[len(res.NavigatepageNums)-1]
	}
	return res
}
