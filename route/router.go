package route

import (
	v1 "tfserver/api/v1"
	"tfserver/middleware"
	"tfserver/util"
	"tfserver/util/log"

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
		//团队模块接口
		routerV1.POST("team/create", middleware.JwtMiddleware(), v1.CreateNewTeam)
		routerV1.POST("team/info/update", middleware.JwtMiddleware(), v1.UpdateTeamInfo)
		routerV1.POST("team/info/get", middleware.JwtMiddleware(), v1.GetTeamInfo)
		routerV1.POST("team/members", middleware.JwtMiddleware(), v1.GetTeamMembers)
		routerV1.POST("team/teams", middleware.JwtMiddleware(), v1.GetUserTeams)
	}

	//注册zap日志框架的中间件
	r.Use(log.GinLogger(), log.GinRecovery(true))

	return r
}
