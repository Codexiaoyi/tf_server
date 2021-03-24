package route

import (
	"tfserver/util"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(util.AppMode)
	r := gin.Default()

	// routerV1 := r.Group("api/v1")
	// {
	// 	//用户模块接口
	// 	routerV1.POST("user/add", v1.AddUser)
	// 	//分类模块接口
	// 	//文章模块接口

	// }

	return r
}
