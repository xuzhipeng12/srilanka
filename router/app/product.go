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

func RegisterProductRouter(api *gin.RouterGroup) {
	user := api.Group("/product")
	{
		user.GET("/list", app.ListProduct)
		user.GET("/listdepts", app.ListAllProductDept)
		user.GET("/:productID", app.GetProductByID)
		user.POST("/add", app.AddProduct)
		user.POST("/update", app.UpdateProduct)
		user.POST("/delete", app.DeleteProduct)
	}
}
