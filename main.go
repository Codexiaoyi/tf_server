package main

import (
	"tfserver/config"
	"tfserver/internal/repository"
	"tfserver/internal/repository/cache"
	"tfserver/internal/route"
	"tfserver/pkg/log"
)

func main() {
	//初始化log
	log.InitLogger()
	//缓存初始化
	cache.InitCache()
	//数据库初始化
	repository.InitDbContext()
	//初始化gin的路由
	ginRoute := route.InitRouter()
	//运行
	ginRoute.Run(config.HttpPort)
}
