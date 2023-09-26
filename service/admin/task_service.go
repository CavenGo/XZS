package admin

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"time"
	"xzs/model"
	"xzs/model/entity"
	"xzs/model/request"
	"xzs/model/response"
	"xzs/pkg/dateutil"
)

func TaskPageListService(req request.TaskPageListRequest) (res response.TaskPageListResponse) {
	tasks := model.TaskExamPageList(req.PageIndex, req.PageSize, req.GradeLevel)
	count := model.TaskExamPageListAllCount(req.GradeLevel)
	page := CalBasePage(req.PageIndex, req.PageSize, int(count), len(tasks))
	res.BasePageResponse = page
	var list []response.TaskPageResponseVM
	for _, task := range tasks {
		temp := response.TaskPageResponseVM{
			Id:             task.Id,
			Title:          task.Title,
			CreateUserName: task.CreateUserName,
			CreateTime:     dateutil.DateFormat(task.CreateTime),
			GradeLevel:     task.GradeLevel,
			Deleted:        task.Deleted,
		}
		list = append(list, temp)
	}
	res.List = list
	return
}

func TaskDeleteService(id int) error {
	return model.TaskExamDeleteById(id)
}

func TaskSelectService(id int) (res response.TaskRequestVM) {
	var paperItems []response.ExamPaperPage
	taskExam := model.TaskSelectById(id)
	copier.Copy(&res, &taskExam)
	textContent, _ := model.TextContentSelectById(taskExam.FrameTextContentId)
	var taskItemObjects []response.TaskItemObject
	json.Unmarshal([]byte(textContent.Content), &taskItemObjects)
	for _, object := range taskItemObjects {
		paperItem := response.ExamPaperPage{}
		examPaper := model.ExamPaperSelectById(object.ExamPaperId)
		copier.Copy(&paperItem, &examPaper)
		paperItem.CreateTime = dateutil.DateFormat(examPaper.CreateTime)
		paperItems = append(paperItems, paperItem)
	}
	res.PaperItems = paperItems
	return
}

func TaskEditService(req response.TaskRequestVM, userName string) (taskExam entity.TaskExam) {
	if req.Id == 0 {
		user, _ := model.FindUserByUserName(userName)
		// 添加
		taskExam = entity.TaskExam{
			Title:          req.Title,
			CreateTime:     time.Now(),
			CreateUser:     user.Id,
			CreateUserName: user.UserName,
			Deleted:        false,
		}

		var taskItemObjects []response.TaskItemObject
		for _, item := range req.PaperItems {
			taskItemObject := response.TaskItemObject{
				ExamPaperId:   item.Id,
				ExamPaperName: item.Name,
			}
			taskItemObjects = append(taskItemObjects, taskItemObject)
		}
		json, _ := json.Marshal(taskItemObjects)
		textContent := entity.TextContent{
			Content:    string(json),
			CreateTime: time.Now(),
		}
		model.TextContentInsertOne(&textContent)

		taskExam.FrameTextContentId = textContent.Id
		model.TaskExamInsertOne(&taskExam)
	} else {
		// 编辑
		taskExam = model.TaskSelectById(req.Id)
		model.TaskExamUpdateById(req.Id, entity.TaskExam{
			Title:      req.Title,
			GradeLevel: req.GradeLevel,
		})

		var taskItemObjects []response.TaskItemObject
		for _, item := range req.PaperItems {
			taskItemObject := response.TaskItemObject{
				ExamPaperId:   item.Id,
				ExamPaperName: item.Name,
			}
			taskItemObjects = append(taskItemObjects, taskItemObject)
		}
		json, _ := json.Marshal(taskItemObjects)
		model.TextContentUpdateContentById(string(json), taskExam.FrameTextContentId)

		model.ClearTaskExamId(taskExam.Id)
	}
	// 更新taskId
	for _, item := range req.PaperItems {
		model.ExamPaperUpdateTaskId(item.Id, taskExam.Id)
	}
	return
}
