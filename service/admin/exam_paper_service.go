package admin

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"strconv"
	"time"
	"xzs/common"
	"xzs/model"
	"xzs/model/entity"
	"xzs/model/request"
	"xzs/model/response"
	"xzs/pkg/dateutil"
)

func ExamPaperPageService(req request.ExamPaperPageRequest) (res response.ExamPaperPageResponse) {
	var list []response.ExamPaperPage
	examPapers := model.ExamPaperPageList(req.PageIndex, req.PageSize, req.Id, req.Level, req.SubjectId, req.PaperType)
	count := model.ExamPaperPageListAllCount(req.Id, req.Level, req.SubjectId, req.PaperType)
	page := CalBasePage(req.PageIndex, req.PageSize, int(count), len(examPapers))
	res.BasePageResponse = page
	for _, v := range examPapers {
		examPaper := response.ExamPaperPage{
			Id:                 v.Id,
			Name:               v.Name,
			QuestionCount:      v.QuestionCount,
			Score:              v.Score,
			CreateTime:         dateutil.DateFormat(v.CreateTime),
			CreateUser:         v.CreateUser,
			SubjectId:          v.SubjectId,
			PaperType:          v.PaperType,
			FrameTextContentId: v.FrameTextContentId,
		}
		list = append(list, examPaper)
	}
	res.List = list
	return
}

func ExamPaperTaskExamPageListService(req request.ExamPaperPageRequest) (res response.ExamPaperPageResponse) {
	var list []response.ExamPaperPage
	examPapers := model.TaskExamPaperPageList(req.PageIndex, req.PageSize, req.Level, req.PaperType)
	count := model.TaskExamPaperPageListAllCount(req.Level, req.PaperType)
	page := CalBasePage(req.PageIndex, req.PageSize, int(count), len(examPapers))
	res.BasePageResponse = page
	for _, v := range examPapers {
		examPaper := response.ExamPaperPage{
			Id:                 v.Id,
			Name:               v.Name,
			QuestionCount:      v.QuestionCount,
			Score:              v.Score,
			CreateTime:         dateutil.DateFormat(v.CreateTime),
			CreateUser:         v.CreateUser,
			SubjectId:          v.SubjectId,
			PaperType:          v.PaperType,
			FrameTextContentId: v.FrameTextContentId,
		}
		list = append(list, examPaper)
	}
	res.List = list
	return
}

func ExamPaperDeleteService(id int) error {
	return model.DeleteExamPaperById(id)
}

func ExamPaperSelectService(id int) (res request.ExamPaperEditRequest, err error) {
	examPaper := model.ExamPaperSelectById(id)
	err = copier.Copy(&res, &examPaper)
	if err != nil {
		return
	}
	res.Level = examPaper.GradeLevel
	// 查询试卷结构
	// [{"name":"计算","questionItems":[{"id":1,"itemOrder":1}]},{"name":"填空题","questionItems":[{"id":2,"itemOrder":2}]}]
	textContent, err := model.TextContentSelectById(examPaper.FrameTextContentId)
	if err != nil {
		return
	}
	var examPaperTitleItemObjects []request.ExamPaperTitleItemObject
	err = json.Unmarshal([]byte(textContent.Content), &examPaperTitleItemObjects)
	if err != nil {
		return
	}

	// 查询所有题目信息
	var questionIds []int
	for _, v := range examPaperTitleItemObjects {
		for _, item := range v.QuestionItems {
			questionIds = append(questionIds, item.Id)
		}
	}
	questions := model.QuestionSelectByIds(questionIds)
	questionMap := make(map[int]entity.Question)
	for _, question := range questions {
		questionMap[question.Id] = question
	}

	// 返回前端的题目结构
	var examPaperTitleItemVMS []request.ExamPaperTitleItemVM
	for _, item := range examPaperTitleItemObjects {
		titleItem := request.ExamPaperTitleItemVM{}
		copier.Copy(&titleItem, &item)
		var questionItems []request.QuestionEditRequest
		for _, questionItem := range item.QuestionItems {
			editRequest := GetQuestionEditRequest(questionMap[questionItem.Id])
			editRequest.ItemOrder = questionItem.ItemOrder
			questionItems = append(questionItems, editRequest)
		}
		titleItem.QuestionItems = questionItems
		examPaperTitleItemVMS = append(examPaperTitleItemVMS, titleItem)
	}
	res.TitleItems = examPaperTitleItemVMS
	res.Score = FormatScore(examPaper.Score)
	if examPaper.PaperType == common.TimeLimit {
		// 限时试卷
		limitDateTime := []string{
			dateutil.DateFormat(examPaper.LimitStartTime),
			dateutil.DateFormat(examPaper.LimitEndTime),
		}
		res.LimitDateTime = limitDateTime
	}
	return
}

func ExamPaperEditService(req request.ExamPaperEditRequest, userName string) (examPaper entity.ExamPaper, err error) {
	user, err := model.FindUserByUserName(userName)
	if err != nil {
		return
	}
	titleItemsVM := req.TitleItems
	frameTextContentList := frameTextContentFromVM(titleItemsVM)
	frameTextContent, _ := json.Marshal(frameTextContentList)
	if req.Id == 0 {
		// 添加
		textContent := entity.TextContent{
			Content:    string(frameTextContent),
			CreateTime: time.Now(),
		}
		err = model.TextContentInsertOne(&textContent)
		if err != nil {
			return
		}
		copier.Copy(&examPaper, &req)
		examPaper.FrameTextContentId = textContent.Id
		examPaper.CreateTime = time.Now()
		examPaper.CreateUser = user.Id
		examPaperFromVM(req, &examPaper, titleItemsVM)
		err = model.ExamPaperInsertOne(&examPaper)
		if err != nil {
			return
		}
	} else {
		// 修改
		examPaper = model.ExamPaperSelectById(req.Id)
		err = model.TextContentUpdateContentById(string(frameTextContent), examPaper.FrameTextContentId)
		if err != nil {
			return
		}
		copier.Copy(&req, &examPaper)
		examPaperFromVM(req, &examPaper, titleItemsVM)
		err = model.ExamPaperUpdateById(examPaper, req.Id)
		if err != nil {
			return
		}
	}
	return
}

func frameTextContentFromVM(titleItemsVM []request.ExamPaperTitleItemVM) []request.ExamPaperTitleItemObject {
	index := 1
	var examPaperTitleItemObjects []request.ExamPaperTitleItemObject
	for _, vm := range titleItemsVM {
		examPaperTitleItemObject := request.ExamPaperTitleItemObject{}
		examPaperTitleItemObject.Name = vm.Name
		var questionItems []request.ExamPaperQuestionItemObject
		for _, v := range vm.QuestionItems {
			item := request.ExamPaperQuestionItemObject{
				Id:        v.Id,
				ItemOrder: index,
			}
			index++
			questionItems = append(questionItems, item)
		}
		examPaperTitleItemObject.QuestionItems = questionItems
		examPaperTitleItemObjects = append(examPaperTitleItemObjects, examPaperTitleItemObject)
	}
	return examPaperTitleItemObjects
}

func examPaperFromVM(examPaperEditRequestVM request.ExamPaperEditRequest, examPaper *entity.ExamPaper, titleItemsVM []request.ExamPaperTitleItemVM) {
	subject, _ := model.SelectSubjectById(examPaperEditRequestVM.SubjectId)
	// 题目总数 总分
	var questionCount int
	var score int
	for _, vm := range titleItemsVM {
		questionCount = questionCount + len(vm.QuestionItems)
		for _, item := range vm.QuestionItems {
			score = score + scoreFromVM(item.Score)
		}
	}
	examPaper.GradeLevel = subject.Level
	examPaper.QuestionCount = questionCount
	examPaper.Score = score
	dateTimes := examPaperEditRequestVM.LimitDateTime
	if examPaper.PaperType == common.TimeLimit {
		start, _ := time.Parse("2006-04-02 15-04-05", dateTimes[0])
		end, _ := time.Parse("2006-04-02 15-04-05", dateTimes[1])
		examPaper.LimitStartTime = start
		examPaper.LimitEndTime = end
	}
}

func scoreFromVM(score string) int {
	if score == "" {
		return 0
	} else {
		float, _ := strconv.ParseFloat(score, 32)
		return int(float * 10)
	}
}
