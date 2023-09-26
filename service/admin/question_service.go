package admin

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
	"time"
	"xzs/common"
	"xzs/model"
	"xzs/model/entity"
	"xzs/model/request"
	"xzs/model/response"
	"xzs/pkg/dateutil"
	"xzs/pkg/htmlutil"
)

func GetQuestionEditRequest(question entity.Question) (res request.QuestionEditRequest) {
	// 题目映射
	textContent, _ := model.TextContentSelectById(question.InfoTextContentId)
	var questionObject request.QuestionObject
	err := json.Unmarshal([]byte(textContent.Content), &questionObject)
	if err != nil {
		return
	}
	err = copier.Copy(&res, &question)
	if err != nil {
		return
	}
	res.Title = questionObject.TitleContent

	// 答案
	switch question.QuestionType {
	case common.SingleChoice:
	case common.TrueFalse:
		res.Correct = question.Correct
	case common.MultipleChoice:
		res.CorrectArray = strings.Split(question.Correct, ",")
	case common.GapFilling:
		var correctContent []string
		for _, object := range questionObject.QuestionItemObjects {
			correctContent = append(correctContent, object.Content)
		}
		res.CorrectArray = correctContent
	case common.ShortAnswer:
		res.Correct = questionObject.Correct
	default:

	}
	res.Score = FormatScore(question.Score)
	res.Analyze = questionObject.Analyze

	// 题目项映射
	var items []request.QuestionEditItemVm
	for _, obj := range questionObject.QuestionItemObjects {
		item := request.QuestionEditItemVm{}
		copier.Copy(&item, &obj)
		items = append(items, item)
	}
	res.Items = items
	return
}

func FormatScore(score int) string {
	if score%10 == 0 {
		return strconv.Itoa(score / 10)
	} else {
		return fmt.Sprintf("%.1f", score/10.0)
	}
}

func QuestionPageListService(req request.QuestionPageRequestVM) (res response.QuestionPageListResponse) {
	questions := model.QuestionPageList(req.PageIndex, req.PageSize, req.Id, req.Level, req.SubjectId, req.QuestionType)
	// 获取分页参数
	count := model.QuestionPageListAllCount(req.Id, req.Level, req.SubjectId, req.QuestionType)
	page := CalBasePage(req.PageIndex, req.PageSize, int(count), len(questions))
	res.BasePageResponse = page
	var list []response.QuestionResponseVM
	for _, question := range questions {
		item := response.QuestionResponseVM{}
		copier.Copy(&item, &question)
		item.CreateTime = dateutil.DateFormat(question.CreateTime)
		item.Score = FormatScore(question.Score)
		textContent, _ := model.TextContentSelectById(question.InfoTextContentId)
		obj := request.QuestionObject{}
		json.Unmarshal([]byte(textContent.Content), &obj)
		item.ShortTitle = htmlutil.Clear(obj.TitleContent)
		list = append(list, item)
	}
	res.List = list
	return
}

func QuestionDeleteService(id int) error {
	return QuestionDeleteService(id)
}

func QuestionSelectService(id int) request.QuestionEditRequest {
	question := model.QuestionSelectById(id)
	return GetQuestionEditRequest(question)
}

func QuestionEditService(req request.QuestionEditRequest, userName string) {
	user, _ := model.FindUserByUserName(userName)
	subject, _ := model.SelectSubjectById(req.SubjectId)
	if req.Id == 0 {
		// 添加
		content := entity.TextContent{
			CreateTime: time.Now(),
		}
		setQuestionInfoFromVM(&content, req)
		model.TextContentInsertOne(&content)
		question := entity.Question{
			SubjectId:         req.SubjectId,
			GradeLevel:        subject.Level,
			CreateTime:        time.Now(),
			QuestionType:      req.QuestionType,
			Status:            1,
			Score:             scoreFromVM(req.Score),
			Difficult:         req.Difficult,
			InfoTextContentId: content.Id,
			Deleted:           false,
			CreateUser:        user.Id,
		}
		setCorrectFromVM(req.Correct, req.CorrectArray, &question)
		model.QuestionInsertOne(&question)
	} else {
		myQuestion := model.QuestionSelectById(req.Id)
		// 修改
		question := entity.Question{
			SubjectId:  req.SubjectId,
			GradeLevel: subject.Level,
			Score:      scoreFromVM(req.Score),
			Difficult:  req.Difficult,
		}
		setCorrectFromVM(req.Correct, req.CorrectArray, &question)
		model.QuestionUpdateById(req.Id, question)

		content := entity.TextContent{}
		setQuestionInfoFromVM(&content, req)
		model.TextContentUpdateContentById(content.Content, myQuestion.InfoTextContentId)
	}
}

func setQuestionInfoFromVM(infoTextContent *entity.TextContent, req request.QuestionEditRequest) {
	var list []request.QuestionEditItemVm
	for _, item := range req.Items {
		obj := request.QuestionEditItemVm{
			Prefix:   item.Prefix,
			Content:  item.Content,
			Score:    scoreFromVM(strconv.Itoa(item.Score)),
			ItemUuid: item.ItemUuid,
		}
		list = append(list, obj)
	}
	questionObj := request.QuestionObject{
		QuestionItemObjects: list,
		Analyze:             req.Analyze,
		TitleContent:        req.Title,
		Correct:             req.Correct,
	}
	content, _ := json.Marshal(questionObj)
	infoTextContent.Content = string(content)
}

func setCorrectFromVM(correct string, correctArray []string, question *entity.Question) {
	if question.QuestionType == common.MultipleChoice {
		question.Correct = strings.Join(correctArray, ",")
	} else {
		question.Correct = correct
	}
}
