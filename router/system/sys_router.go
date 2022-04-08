/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/2/26 6:40 下午
 **/

package system

import (
	"github.com/gin-gonic/gin"
	"srilanka/service/system"
)

func RegisterUserRouter(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.GET("/list", system.ListUser)
		user.POST("/add", system.AddUser)
		user.POST("/update", system.UpdateUser)
		user.POST("/delete", system.DeleteUser)
		user.GET("/info", system.GetUserInfo)
	}
}

func RegisterSysRouter(api *gin.RouterGroup) {
	sys := api.Group("")
	{
		sys.POST("/login", system.Login)
		sys.POST("/logout", system.Logout)
	}
}
