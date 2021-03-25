package route

import (
	v1 "tfserver/api/v1"
	"tfserver/util"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(util.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		//用户模块接口
		routerV1.POST("user/add", v1.AddUser)
	}

	return r
}