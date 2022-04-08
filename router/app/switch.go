/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/3/12 6:24 下午
 **/
package app

import (
	"github.com/gin-gonic/gin"
	"srilanka/service/app"
)

func RegisterSwitchTemplateRouter(api *gin.RouterGroup) {
	user := api.Group("/switch/template")
	{
		user.GET("/list", app.ListTemplate)
		user.GET("/:templateID", app.GetTemplateByID)
		user.POST("/add", app.AddTemplate)
		user.POST("/update", app.UpdateTemplate)
		user.POST("/delete", app.DeleteTemplate)
	}
}
func RegisterSwitchRouter(api *gin.RouterGroup) {
	user := api.Group("/switch")
	{
		user.GET("/list", app.ListSwitch)
		user.POST("/add", app.AddSwitch)
		user.POST("/update", app.UpdateSwitch)
		user.POST("/delete", app.DeleteSwitch)
	}
}

func RegisterSwitchPublicRouter(api *gin.RouterGroup) {
	sw := api.Group("/switch")
	{
		sw.GET("/share", app.GetSharedSwitchByID)
	}
}
