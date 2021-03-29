package route

import (
	v1 "tfserver/api/v1"
	"tfserver/middleware"
	"tfserver/util"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(util.AppMode)
	r := gin.Default()
	routerV1 := r.Group("api/v1")
	{
		//账号模块
		routerV1.POST("account/register", v1.Register)
		routerV1.POST("account/login", v1.Login)
		//用户模块接口
		routerV1.POST("user/info/get", middleware.JwtMiddleware(), v1.GetUserInfo)
		routerV1.POST("user/info/update", middleware.JwtMiddleware(), v1.UpdateUserInfo)
	}

	return r
}
