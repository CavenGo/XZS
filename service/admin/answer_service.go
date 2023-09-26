package admin

import (
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"xzs/model"
	"xzs/model/request"
	"xzs/model/response"
	"xzs/pkg/dateutil"
)

func AnswerPageJudgeListService(req request.ExamPaperAnswerPageRequestVM) (res response.PageJudgeListResponse) {
	answers := model.ExamPaperAnswerPageList(req.PageIndex, req.PageSize, req.SubjectId)
	count := model.ExamPaperAnswerPageListAllCount(req.SubjectId)
	page := CalBasePage(req.PageIndex, req.PageSize, int(count), len(answers))

	res.BasePageResponse = page

	var list []response.ExamPaperAnswerPageResponseVM
	for _, answer := range answers {
		item := response.ExamPaperAnswerPageResponseVM{}
		copier.Copy(&item, &answer)
		subject, _ := model.SelectSubjectById(answer.SubjectId)
		item.SubjectName = subject.Name
		item.CreateTime = dateutil.DateFormat(answer.CreateTime)
		user, _ := model.GetUserById(answer.CreateUser)
		item.UserName = user.UserName
		item.SystemScore = scoreToVM(answer.SystemScore)
		item.UserScore = scoreToVM(answer.UserScore)
		item.PaperScore = scoreToVM(answer.PaperScore)
		item.DoTime = secondToVM(answer.DoTime)
		list = append(list, item)
	}
	res.List = list
	return
}

func scoreToVM(score int) string {
	if score%10 == 0 {
		return strconv.Itoa(score / 10)
	} else {
		return fmt.Sprintf("%.1f", score/1.0)
	}
}

func secondToVM(second int) string {
	dateTimes := ""
	days := second / (60 * 60 * 24)
	hours := (second % (60 * 60 * 24)) / (60 * 60)
	minutes := (second % (60 * 60)) / 60
	seconds := second % 10
	if days > 0 {
		dateTimes = fmt.Sprintf("%d天%d时%d分%d秒", days, hours, minutes, seconds)
	} else if hours > 0 {
		dateTimes = fmt.Sprintf("%d时%d分%d秒", hours, minutes, seconds)
	} else if minutes > 0 {
		dateTimes = fmt.Sprintf("%d分%d秒", minutes, seconds)
	} else {
		dateTimes = fmt.Sprintf("%d秒", seconds)
	}

	return dateTimes
}
