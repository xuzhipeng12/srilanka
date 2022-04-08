package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
	_ "srilanka/docs"
	"srilanka/global"
	"srilanka/models"
	router2 "srilanka/router"
	"srilanka/tools"
)

var swagger = flag.Bool("swagger", true, "是否启动swagger接口文档,默认不启动")
var configFile = flag.String("configFile", "config/config.yml", "配置文件路径")
var initDb = flag.Bool("initDb", false, "是否初始化数据库")

func init() {
	flag.Parse()
	tools.ConfigRead(*configFile)
	tools.LogInit()
}

//@title gin示例 API
//@version 0.0.1
//@description  相关接口文档
//@host 127.0.0.1:8080
//@BasePath
func main() {
	global.DBSetup()
	if *initDb {
		err := models.InitDbModels()
		if err == nil {
			fmt.Println("db init succeed")
			os.Exit(0)
		} else {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	}
	tools.Log.Info("start server...")
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	// gin路由初始化
	r := router2.Initrouter()
	if *swagger {
		//启动访问swagger文档
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	tools.Log.Info("listen on :%s", tools.Cfg.ListenPort)
	//监听端口
	r.Run(":" + tools.Cfg.ListenPort)
}
