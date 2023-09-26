package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"go_code/xzs/common"
	"go_code/xzs/config"
	"go_code/xzs/controller"
	"go_code/xzs/controller/admin"
	"go_code/xzs/middleware"
)

//web框架
func SetupRouters() *gin.Engine {

	// 如果是dev环境开启debug模式
	if config.GlobalConf.Server.Env != "pro" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(requestid.New())
	// 单独调试前端时需要配置cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:8000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "request-ajax"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(config.GinLogger(), config.GinRecovery(true))

	r.StaticFS("/admin", http.Dir("static/admin"))
	r.StaticFS("/student", http.Dir("static/student"))

	// 后台测试接口
	r.POST("/api/admin/test", admin.TestApi.Test)

	// 启用session
	r.Use(sessions.Sessions("xzsSession", cookie.NewStore([]byte("secret"))))

	// 不需要登录的路由
	r.POST("/api/user/login", controller.UserApi.Login)

	// 中间件
	r.Use(middleware.AuthLogin())

	// 公共的接口
	r.POST("/api/user/logout", controller.UserApi.Logout)

	// 后台路由
	adminGroup := r.Group("/api/admin")
	{
		adminGroup.POST("/dashboard/index", admin.DashoboardApi.Index)

		adminGroup.POST("/user/page/list", admin.UserApi.PageList)
		adminGroup.POST("/user/changeStatus/:id", admin.UserApi.ChangeStatus)
		adminGroup.POST("/user/current", admin.UserApi.Current)
		adminGroup.POST("/user/select/:id", admin.UserApi.Select)
		adminGroup.POST("/user/edit", admin.UserApi.Edit)
		adminGroup.POST("/user/update", admin.UserApi.Update)
		adminGroup.POST("/user/delete/:id", admin.UserApi.Delete)
		adminGroup.POST("/user/selectByUserName", admin.UserApi.SelectByUserName)
		adminGroup.POST("/user/event/page/list", admin.UserApi.EventPageList)

		adminGroup.POST("/education/subject/list", admin.EducationApi.List)
		adminGroup.POST("/education/subject/page", admin.EducationApi.PageList)
		adminGroup.POST("/education/subject/edit", admin.EducationApi.Edit)
		adminGroup.POST("/education/subject/select/:id", admin.EducationApi.Select)
		adminGroup.POST("/education/subject/delete/:id", admin.EducationApi.Delete)

		adminGroup.POST("/exam/paper/page", admin.ExamPaperApi.Page)
		adminGroup.POST("/exam/paper//taskExamPage", admin.ExamPaperApi.TaskExamPageList)
		adminGroup.POST("/exam/paper/delete/:id", admin.ExamPaperApi.Delete)
		adminGroup.POST("/exam/paper/select/:id", admin.ExamPaperApi.Select)
		adminGroup.POST("/exam/paper/edit", admin.ExamPaperApi.Edit)

		adminGroup.POST("/question/page", admin.QuestionApi.PageList)
		adminGroup.POST("/question/select/:id", admin.QuestionApi.Select)
		adminGroup.POST("/question/delete/:id", admin.QuestionApi.Delete)
		adminGroup.POST("/question/edit", admin.QuestionApi.Edit)

		adminGroup.POST("/task/page", admin.TaskApi.PageList)
		adminGroup.POST("/task/select/:id", admin.TaskApi.Select)
		adminGroup.POST("/task/delete/:id", admin.TaskApi.Delete)
		adminGroup.POST("/task/edit", admin.TaskApi.Edit)

		adminGroup.POST("/examPaperAnswer/page", admin.AnswerApi.PageJudgeList)

		adminGroup.POST("/message/send", admin.MessageApi.Send)
		adminGroup.POST("/message/page", admin.MessageApi.PageList)

	}

	// 访问未定义的路由时的提示
	r.NoRoute(func(c *gin.Context) {
		common.ResponseFailWithCodeMsg(c, common.InnerError, "Not Found (#404)")
		return
	})

	return r
}
