package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"xzs/model"
	"xzs/model/response"
	"xzs/pkg/dateutil"
)

func DashboardIndexService(c *gin.Context) response.DashBoardIndexResponse {
	var res response.DashBoardIndexResponse

	examPaperAllCount := model.ExamPaperAllCount(c)
	examPaperAnswerAllCount := model.ExamPaperAnswerAllCount()
	examPaperQuestionCustomerAnswerAllCount := model.ExamPaperQuestionCustomerAnswerAllCount()
	questionAllCount := model.QuestionAllCount()

	res.ExamPaperCount = examPaperAllCount
	res.QuestionCount = questionAllCount
	res.DoExamPaperCount = examPaperAnswerAllCount
	res.DoQuestionCount = examPaperQuestionCustomerAnswerAllCount

	res.MothDayUserActionValue = getMothDayUserActionValue()
	res.MothDayText = dateutil.MothDay()
	res.MothDayDoExamQuestionValue = getMothDayDoExamQuestionValue()

	return res
}

func getMothDayUserActionValue() []int {
	var res []int
	eventLogSlice := model.SelectEventLogByStartEnd(dateutil.GetMonthStartDay(), dateutil.GetMonthEndDay())
	eventLogMap := make(map[string]int)
	for _, v := range eventLogSlice {
		eventLogMap[v.Name] = v.Value
	}
	// 本月的天数
	days := time.Now().Day()
	yearMonthStr := time.Now().Format("2006-01")
	for i := 1; i <= days; i++ {
		value, ok := eventLogMap[yearMonthStr+"-"+fmt.Sprintf("%02d", i)]
		if ok {
			res = append(res, value)
		} else {
			res = append(res, 0)
		}
	}
	return res
}

func getMothDayDoExamQuestionValue() []int {
	var res []int
	CustomerAnswerSlice := model.SelectCustomerAnswerByStartEnd(dateutil.GetMonthStartDay(), dateutil.GetMonthEndDay())
	customerAnswerMap := make(map[string]int)
	for _, v := range CustomerAnswerSlice {
		customerAnswerMap[v.Name] = v.Value
	}
	// 本月的天数
	days := time.Now().Day()
	yearMonthStr := time.Now().Format("2006-01")
	for i := 1; i <= days; i++ {
		value, ok := customerAnswerMap[yearMonthStr+"-"+fmt.Sprintf("%02d", i)]
		if ok {
			res = append(res, value)
		} else {
			res = append(res, 0)
		}
	}
	return res
}
