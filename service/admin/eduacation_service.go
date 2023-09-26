package admin

import (
	"xzs/model"
	"xzs/model/entity"
	"xzs/model/request"
	"xzs/model/response"
)

func EducationListService() []entity.Subject {
	subject := model.SelectAllSubject()
	return subject
}

func EducationPageListService(req request.EducationPageListRequest) (res response.EducationPageListResponse) {
	var list []response.PageListSubject
	subjects := model.SubjectPageList(req.PageIndex, req.PageSize, req.Id, req.Level)
	count := model.SubjectAllCountByUserNameRole(req.Id, req.Level)
	// 计算分页
	page := CalBasePage(req.PageIndex, req.PageSize, int(count), len(subjects))
	res.BasePageResponse = page
	for _, v := range subjects {
		subject := response.PageListSubject{
			Id:        v.Id,
			Name:      v.Name,
			Level:     v.Level,
			LevelName: v.LevelName,
		}
		list = append(list, subject)
	}
	res.List = list
	return
}

func EducationEditService(req request.EducationEditRequest) (err error) {
	subject := entity.Subject{
		Name:      req.Name,
		Level:     req.Level,
		LevelName: req.LevelName,
	}
	if req.Id != 0 {
		// 更新
		err = model.UpdateSubjectById(req.Id, subject)
	} else {
		// 插入
		err = model.AddSubject(&subject)
	}
	return err
}

func EducationSelectService(id int) (res response.PageListSubject, err error) {
	subject, err := model.SelectSubjectById(id)
	if err != nil {
		return
	}
	res.Id = subject.Id
	res.Name = subject.Name
	res.Level = subject.Level
	res.LevelName = subject.LevelName
	return
}

func EducationDeleteService(id int) error {
	return model.DeleteSubjectById(id)
}
