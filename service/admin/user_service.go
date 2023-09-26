package admin

import (
	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
	"time"
	"xzs/common"
	"xzs/model"
	"xzs/model/entity"
	"xzs/model/request"
	"xzs/model/response"
	"xzs/pkg/dateutil"
	"xzs/pkg/encryptutil"
)

func PageListService(req request.PageListRequest) (res response.PageListResponse) {
	var list []response.PageListUser
	// 分页查询数据
	users := model.UserPageList(req.PageIndex, req.PageSize, req.Role, req.UserName)
	// 获取总数
	count := model.UserAllCountByUserNameRole(req.UserName, req.Role)
	// 计算总页数
	pages := count / int64(req.PageSize)
	if count%int64(req.PageSize) != 0 {
		pages = pages + 1
	}
	// 计算分页
	page := CalBasePage(req.PageIndex, req.PageSize, int(count), len(users))
	res.BasePageResponse = page
	for _, v := range users {
		pageListUser := response.PageListUser{
			Id:             v.Id,
			UserName:       v.UserName,
			UserUuid:       v.UserUuid,
			RealName:       v.RealName,
			Age:            v.Age,
			Role:           v.Role,
			Sex:            v.Sex,
			BirthDay:       dateutil.DateFormat(v.BirthDay),
			Phone:          v.Phone,
			LastActiveTime: dateutil.DateFormat(v.LastActiveTime),
			CreateTime:     dateutil.DateFormat(v.CreateTime),
			ModifyTime:     dateutil.DateFormat(v.ModifyTime),
			Status:         v.Status,
			UserLevel:      v.UserLevel,
			ImagePath:      v.ImagePath,
		}
		list = append(list, pageListUser)
	}
	res.List = list
	return res
}

func UserChangeStatus(id int) (newStatus int, err error) {
	user, err := model.GetUserById(id)
	if err != nil {
		return
	}
	userStatus := common.UserEnable
	if user.Status == common.UserEnable {
		userStatus = common.UserDisable
	}
	err = model.UpdateStatusById(id, userStatus)
	return userStatus, err
}

func UserCurrentService(userName string) (res response.PageListUser, err error) {
	user, err := model.FindUserByUserName(userName)
	if err != nil {
		return
	}
	userMapVm(user, &res)
	return
}

func UserSelectService(id int) (res response.PageListUser, err error) {
	user, err := model.GetUserById(id)
	if err != nil {
		return
	}
	userMapVm(user, &res)
	return
}

func userMapVm(user entity.User, res *response.PageListUser) {
	res.Id = user.Id
	res.Age = user.Age
	res.Sex = user.Sex
	res.BirthDay = dateutil.DateFormat(user.BirthDay)
	res.Phone = user.Phone
	res.ImagePath = user.ImagePath
	res.CreateTime = dateutil.DateFormat(user.CreateTime)
	res.LastActiveTime = dateutil.DateFormat(user.LastActiveTime)
	res.UserLevel = user.UserLevel
	res.UserName = user.UserName
	res.UserUuid = user.UserUuid
	res.ModifyTime = dateutil.DateFormat(user.ModifyTime)
	res.Role = user.Role
	res.RealName = user.RealName
	res.Status = user.Status
}

func UserEditService(req request.UserEditRequest) (res common.RestResponse) {
	res.Code = common.Ok
	res.Message = common.Ok.Msg()
	if req.Id == 0 {
		// 插入
		existUser, err := model.FindUserByUserName(req.UserName)
		if err == nil && existUser.Id != 0 {
			res.Code = 2
			res.Message = "用户已存在"
			return
		}
		if req.Password == "" {
			res.Code = 3
			res.Message = "密码不能为空"
			return
		}

	}

	user := entity.User{
		UserName:  req.UserName,
		RealName:  req.RealName,
		Age:       req.Age,
		Status:    req.Status,
		Sex:       req.Sex,
		Phone:     req.Phone,
		Role:      req.Role,
		UserLevel: req.UserLevel,
	}
	if req.BirthDay != "" {
		parse, err := time.Parse("2006-01-02", req.BirthDay)
		if err == nil {
			user.BirthDay = parse
		}
	}
	if req.Password != "" {
		encodePwd, err := encryptutil.RsaEncode(req.Password)
		if err != nil {
			res.Code = common.InnerError
			res.Message = "密码生成失败"
			return
		}
		user.Password = encodePwd
	}
	if req.Id == 0 {
		user.CreateTime = time.Now()
		user.LastActiveTime = time.Now()
		user.Deleted = false
		user.UserUuid = uuid.NewV4().String()
		err := model.AddUser(&user)
		if err != nil {
			res.Code = common.InnerError
			res.Message = "添加用户失败"
			return
		}
	} else {
		user.ModifyTime = time.Now()
		err := model.UpdateUserById(req.Id, user)
		if err != nil {
			if err != nil {
				res.Code = common.InnerError
				res.Message = "更新用户信息失败"
				return
			}
		}
	}
	return
}

func UserSelectByUserNameService(userName string) []entity.UserKeyValue {
	list := model.SelectUserKeyValueByUserName(userName)
	return list
}

func EventPageListService(req request.UserEventPageRequestVM) (res response.UserEventLogResponse) {
	var list []response.UserEventLogVM
	events := model.UserEventLogPageList(req.PageIndex, req.PageSize, req.UserId, req.UserName)
	count := model.UserEventLogPageListAllCount(req.UserId, req.UserName)
	page := CalBasePage(req.PageIndex, req.PageSize, int(count), len(events))
	res.BasePageResponse = page
	for _, v := range events {
		tmp := response.UserEventLogVM{}
		copier.Copy(&tmp, v)
		tmp.CreateTime = dateutil.DateFormat(v.CreateTime)
		list = append(list, tmp)
	}
	res.List = list
	return
}
/*以上代码是一个 Go 语言编写的后端服务程序，主要包含了用户管理和事件管理的功能。下面逐个函数进行解释：

PageListService(req request.PageListRequest) (res response.PageListResponse)
该函数是用户列表分页查询的功能实现。通过传入的 PageListRequest 对象的参数，查询用户列表，并将查询结果返回到 PageListResponse 对象中。其中，PageListUser 对象是用户列表的每一项数据，包含了用户的各种信息。

UserChangeStatus(id int) (newStatus int, err error)
该函数是更改用户状态的功能实现。通过传入的用户 ID 查询用户信息，然后更改用户的状态（启用或禁用），最后将更改后的状态返回。如果查询用户信息出错，则返回错误信息。

UserCurrentService(userName string) (res response.PageListUser, err error)
该函数是查询当前登录用户的信息的功能实现。通过传入的用户名参数，查询用户信息，并将查询结果返回到 PageListUser 对象中。如果查询用户信息出错，则返回错误信息。

UserSelectService(id int) (res response.PageListUser, err error)
该函数是查询用户信息的功能实现。通过传入的用户 ID 参数，查询用户信息，并将查询结果返回到 PageListUser 对象中。如果查询用户信息出错，则返回错误信息。

userMapVm(user entity.User, res *response.PageListUser)
该函数是将 entity.User 对象映射到 response.PageListUser 对象的功能实现。在用户信息查询的过程中，将 entity.User 对象转换为 response.PageListUser 对象，以便返回用户列表的每一项数据。

UserEditService(req request.UserEditRequest) (res common.RestResponse)
该函数是添加或更新用户信息的功能实现。通过传入的 UserEditRequest 对象，判断是添加用户还是更新用户信息，并根据传入的参数对用户信息进行修改。如果是添加用户，则需要对密码进行加密，同时生成一个 UUID 作为用户的唯一标识。如果是更新用户信息，则需要更新 ModifyTime 字段。最后，返回一个 RestResponse 对象，包含了修改结果的状态码和提示信息。

UserSelectByUserNameService(userName string) []entity.UserKeyValue
该函数是根据用户名查询用户信息的功能实现。通过传入的用户名参数，查询用户信息，并返回一个 UserKeyValue 对象的数组，其中包含了用户的 ID 和用户名。

EventPageListService(req request.UserEventPageRequestVM) (res response.UserEventLogResponse)
该函数是用户事件列表分页查询的功能实现。通过传入的 UserEventPageRequestVM 对象的参数，查询用户事件列表，并将查询结果返回到 UserEventLogResponse 对象中。其中，UserEventLogVM 对象是用户事件列表的每一项数据，包含了事件的各种信息。

以上就是对该程序中各个函数的详细解释。*/