package main

import (
	"fmt"
	_"github.com/gin-gonic/gin"
	"context"
	"github.com/json-iterator/go/extra"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	_"strconv"
	"time"
	"go_code/xzs/config"
	"go_code/xzs/global"
	"go_code/xzs/model"
	"go_code/xzs/routers"
)

func init(){
	// jsoniter开启php兼容模式
	// 注意构建时还需要添加构建参数-tags=jsoniter
	extra.RegisterFuzzyDecoders()
}

func main(){

	// 读取配置文件
	config.InitSetting()
	// 初始化zap日志组件
	err := config.InitZap()
	if err != nil {
		log.Fatalln("初始化zap日志组件失败，", err)
	}
	// 初始化数据库
	global.Db, err = model.InitMysql()
	if err != nil {
		log.Fatalln("初始化数据库失败，", err)
	}else{
		fmt.Println("初始化数据库成功")
	}

	// 初始化协程池
	err = config.InitPool()
	if err != nil {
		log.Fatalln("初始化协程池出错，", err)
	}

	// 启动服务
	runServe()

}

func runServe() {
	r := routers.SetupRouters()
	// github.com/gin-contrib/pprof 性能分析
	// pprof.Register(r)
	srv := &http.Server{
		//Addr:    ":" + strconv.FormatInt(config.GlobalConf.Server.Port, 10),
		Addr: ":8000",
		Handler: r,
	}
	go func() {
		// 服务连接
		err := srv.ListenAndServe(); 
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown:", zap.Error(err))
	}
	zap.L().Info("Server exiting")


}
