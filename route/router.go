package route

import (
	v1 "tfserver/api/v1"
	"tfserver/config"
	"tfserver/middleware"
	"tfserver/util/log"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.AppMode)
	r := gin.Default()
	routerV1 := r.Group("api/v1")
	{
		//账号模块
		routerV1.POST("account/register", v1.Register)
		routerV1.POST("account/login", v1.Login)
		//用户模块接口
		routerV1.POST("user/info/get", middleware.JwtMiddleware(), v1.GetUserInfo)
		routerV1.POST("user/info/update", middleware.JwtMiddleware(), v1.UpdateUserInfo)
		routerV1.POST("user/avatar/upload", middleware.JwtMiddleware(), v1.UploadUserAvatar)
		routerV1.POST("user/avatar/get", middleware.JwtMiddleware(), v1.GetUserAvatar)
		//用户相册模块接口
		routerV1.POST("user/album/create", middleware.JwtMiddleware(), v1.CreateNewUserAlbum)
		routerV1.POST("user/album/albums", middleware.JwtMiddleware(), v1.GetUserAlbums)
		routerV1.POST("user/album/info/get", middleware.JwtMiddleware(), v1.GetUserAlbumInfo)
		routerV1.POST("user/album/info/update", middleware.JwtMiddleware(), v1.UpdateUserAlbumInfo)
		routerV1.POST("user/album/media/medias", middleware.JwtMiddleware(), v1.GetUserAlbumMedias)
		routerV1.POST("user/album/thumbnail/get", middleware.JwtMiddleware(), v1.GetThumbnail)
		routerV1.POST("user/album/cover/set", middleware.JwtMiddleware(), v1.SetUserAlbumCover)
		routerV1.POST("user/album/upload/key/get", middleware.JwtMiddleware(), v1.GetUploadKey)
		routerV1.POST("user/album/upload/media/success", middleware.JwtMiddleware(), v1.UploadMedia)
		//团队模块接口
		routerV1.POST("team/create", middleware.JwtMiddleware(), v1.CreateNewTeam)
		routerV1.POST("team/info/update", middleware.JwtMiddleware(), v1.UpdateTeamInfo)
		routerV1.POST("team/info/get", middleware.JwtMiddleware(), v1.GetTeamInfo)
		routerV1.POST("team/members", middleware.JwtMiddleware(), v1.GetTeamMembers)
		routerV1.POST("team/teams", middleware.JwtMiddleware(), v1.GetUserTeams)
		routerV1.POST("team/member/add", middleware.JwtMiddleware(), v1.AddMember)
		routerV1.POST("team/member/leave", middleware.JwtMiddleware(), v1.MemberLeave)
		routerV1.POST("team/member/remove", middleware.JwtMiddleware(), v1.RemoveMember)
		routerV1.POST("team/leader/transform", middleware.JwtMiddleware(), v1.TransformLeader)
	}

	//注册zap日志框架的中间件
	r.Use(log.GinLogger(), log.GinRecovery(true))

	return r
}
