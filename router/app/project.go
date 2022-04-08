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

func RegisterProjectRouter(api *gin.RouterGroup) {
	user := api.Group("/project")
	{
		user.GET("/list", app.ListProject)
		user.GET("/:projectID", app.GetProjectByID)
		user.POST("/add", app.AddProject)
		user.POST("/update", app.UpdateProject)
		user.POST("/delete", app.DeleteProject)
	}
}
