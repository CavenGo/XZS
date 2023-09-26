package admin

import (
	"errors"
	"strings"
	"time"
	"xzs/model"
	"xzs/model/entity"
	"xzs/model/request"
	"xzs/model/response"
	"xzs/pkg/dateutil"
)

func MessageSend(req request.MessageSendVM, userName string) error {
	user, err := model.FindUserByUserName(userName)
	if err != nil {
		return errors.New("登录失效请重新登录")
	}
	userList := model.FindUserByIds(req.ReceiveUserIds)
	message := entity.Message{
		Title:            req.Title,
		Content:          req.Content,
		ReadCount:        0,
		ReceiveUserCount: len(req.ReceiveUserIds),
		CreateTime:       time.Now(),
		SendRealName:     user.RealName,
		SendUserID:       user.Id,
		SendUserName:     user.UserName,
	}
	err = model.MessageInsert(&message)
	if err != nil {
		return errors.New("添加message失败")
	}
	messageUsers := make([]entity.MessageUser, len(userList))
	for k, v := range userList {
		messageUsers[k] = entity.MessageUser{
			CreateTime:      time.Now(),
			Readed:          0,
			ReceiveRealName: v.RealName,
			ReceiveUserID:   v.Id,
			ReceiveUserName: v.UserName,
			MessageID:       message.ID,
		}
	}
	err = model.MessageUserBatchSave(&messageUsers)
	if err != nil {
		return errors.New("添加messageUser失败")
	}
	return nil
}

func MessagePageListService(req request.MessagePageRequestVM) (res response.MessagePageListResponse) {
	var list []response.MessageResponseVM
	messages := model.MessagePaperPageList(req.PageIndex, req.PageSize, req.SendUserName)
	count := model.MessagePageListAllCount(req.SendUserName)
	page := CalBasePage(req.PageIndex, req.PageSize, int(count), len(messages))
	res.BasePageResponse = page
	for _, v := range messages {
		findList := model.MessageUserFindList(map[string]interface{}{
			"message_id": v.ID,
		})
		receives := ""
		if len(findList) > 0 {
			receivesList := make([]string, 0)
			for _, v := range findList {
				receivesList = append(receivesList, v.ReceiveUserName)
			}
			receives = strings.Join(receivesList, ",")
		}

		tmp := response.MessageResponseVM{
			Id:           v.ID,
			Title:        v.Title,
			Content:      v.Content,
			CreateTime:   dateutil.DateFormat(v.CreateTime),
			SendUserName: v.SendUserName,
			ReadCount:    len(findList),
			Receives:     receives,
		}
		list = append(list, tmp)
	}
	res.List = list
	return
}
