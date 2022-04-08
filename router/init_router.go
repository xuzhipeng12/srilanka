/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 6:40 下午
 **/
package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	md "srilanka/middleware"
	"srilanka/router/app"
	"srilanka/router/system"
)

var (
	router *gin.Engine
)

func Initrouter() *gin.Engine {
	// 禁用控制台颜色
	//gin.DisableConsoleColor()

	//gin.New()返回一个*Engine 指针
	//而gin.Default()不但返回一个*Engine 指针，而且还进行了debugPrintWARNINGDefault()和engine.Use(Logger(), Recovery())其他的一些中间件操作
	router = gin.Default()
	//router = gin.New()

	//使用日志
	//router.Use(gin.Logger())
	//使用Panic处理方案
	//router.Use(gin.Recovery())

	router.Use(md.InitErrorHandler)
	router.Use(md.InitAccessLogMiddleware)

	// 未知调用方式
	router.NoMethod(md.InitNoMethodJson)
	// 未知路由处理
	router.NoRoute(md.InitNoRouteJson)

	// 允许跨域
	router.Use(md.Cors)
	// token 校验
	api := router.Group("/api")
	app.RegisterSwitchPublicRouter(api)
	noAuthApi := router.Group("/api/sys")
	api.Use(md.JWTAuth())
	system.RegisterSysRouter(noAuthApi)
	system.RegisterUserRouter(api)
	app.RegisterProductRouter(api)
	app.RegisterProjectRouter(api)
	app.RegisterVersionRouter(api)
	app.RegisterSwitchTemplateRouter(api)
	app.RegisterSwitchRouter(api)

	router.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "post",
		})
	})

	return router

}
