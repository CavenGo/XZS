package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
	"xzs/common"
	"xzs/config"
	"xzs/model"
	"xzs/model/entity"
	"xzs/model/request"
	"xzs/pkg/encryptutil"
	"xzs/pkg/jwtutil"
)

var CookieExpireDuration int = 3600 * 24 * 7

func LoginService(c *gin.Context, req request.LoginRequest) common.RestResponse {
	res := common.RestResponse{
		Code:    common.Ok,
		Message: common.Ok.Msg(),
	}
	// 登录逻辑
	user, err := model.FindUserByUserName(req.UserName)
	if err != nil {
		res.Code = common.AuthError
		res.Message = common.AuthError.Msg()
		return res
	}
	// 判断密码是否正确
	pwd, err := encryptutil.RsaDecode(user.Password)
	if err != nil || req.Password != pwd {
		res.Code = common.AuthError
		res.Message = common.AuthError.Msg()
		return res
	}
	// 判断用户状态
	if user.Status == 2 {
		res.Code = common.AuthError
		res.Message = "用户被禁用"
		return res
	}

	// 登录成功设置session
	session := sessions.Default(c)
	session.Set("userName", req.UserName)
	err = session.Save()
	if err != nil {
		res.Code = common.AuthError
		res.Message = "设置登录状态错误"
		return res
	}
	// 判断是否需要设置cookie
	if req.Remember {
		token, _, err := jwtutil.GenToken(req.UserName)
		if err != nil {
			res.Code = common.AuthError
			res.Message = "记住密码设置失败"
			return res
		}
		c.SetCookie("token", token, CookieExpireDuration, "/", "localhost", false, true)
	}
	userEventLog := entity.UserEventLog{
		UserId:     user.Id,
		UserName:   user.UserName,
		RealName:   user.RealName,
		CreateTime: time.Now(),
		Content:    user.UserName + "  登录了学之思考试系统",
	}
	// 添加登录日志
	config.GlobalPool.Submit(func() {
		model.AddUserEventLog(&userEventLog)
	})
	newUser := entity.User{
		UserName:  user.UserName,
		ImagePath: user.ImagePath,
	}
	res.Response = newUser
	return res
}

func LogoutService(c *gin.Context) {
	// 清除session
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	// 清除cookie
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	userName := c.GetString("userName")
	// 记录日志
	config.GlobalPool.Submit(func() {
		if userName == "" {
			return
		}
		user, err := model.FindUserByUserName(userName)
		if err != nil {
			return
		}
		userEventLog := entity.UserEventLog{
			UserId:     user.Id,
			UserName:   user.UserName,
			RealName:   user.RealName,
			CreateTime: time.Now(),
			Content:    user.UserName + "  登出了学之思开源考试系统",
		}
		err = model.AddUserEventLog(&userEventLog)
	})
}

func UserUpdateService(userName string, req request.UserUpdateRequest) (err error) {
	user := entity.User{
		RealName:   req.RealName,
		Phone:      req.Phone,
		ModifyTime: time.Now(),
	}
	err = model.UpdateUserByUserName(userName, user)
	return
}

func UserDeleteService(id int) (err error) {
	return model.DeleteUserById(id)
}
/*这段代码定义了一些服务函数，用于处理与用户认证、用户信息更新和用户删除相关的操作。下面是对每个函数的详细解释：

LoginService
该函数用于处理用户登录请求。传入参数为 gin.Context 和一个 request.LoginRequest 类型的结构体。该函数首先会根据传入的用户名在数据库中查找对应的用户，如果找不到则返回认证错误的响应；如果找到了，则会将请求中的密码进行解密后与数据库中的密码进行比较，如果不匹配则返回认证错误的响应；如果匹配但用户状态为禁用，则返回错误的响应。如果登录成功，该函数会将用户信息保存到 session 中，如果请求中需要记住密码，则还会将一个 token 存储到 cookie 中。最后，该函数会返回一个包含用户信息的响应。

LogoutService
该函数用于处理用户登出请求。该函数会清除 session 和 cookie 中保存的登录信息，并记录用户登出的日志。

UserUpdateService
该函数用于处理用户信息更新请求。传入参数为用户名和一个 request.UserUpdateRequest 类型的结构体。该函数会将传入的结构体中的信息更新到数据库中对应的用户信息中。

UserDeleteService
该函数用于处理用户删除请求。传入参数为用户 ID，该函数会在数据库中删除对应的用户信息。

在这些服务函数中，使用了一些其他的模块和库，例如：

gin：一个基于 Go 语言的 Web 框架，用于处理 HTTP 请求和响应。
sessions：一个 Gin 中的中间件，用于处理 session 相关的操作。
jwtutil：一个用于生成和解析 JWT 的库。
encryptutil：一个用于加密和解密数据的库。
model：一个用于与数据库交互的模块。
config：一个用于读取配置文件的模块。
entity：一个用于定义数据库中的实体类的模块。
common：一个用于定义一些常用的响应状态码和消息的模块。*/