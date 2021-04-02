package main

import (
	"tfserver/config"
	"tfserver/repository"
	"tfserver/route"
	"tfserver/util/log"
)

func main() {
	//数据库初始化
	repository.InitDbContext()
	//初始化log
	log.InitLogger()
	//初始化gin的路由
	ginRoute := route.InitRouter()
	//运行
	ginRoute.Run(config.HttpPort)
}
