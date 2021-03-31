package main

import (
	"tfserver/repository"
	"tfserver/route"
	"tfserver/util"
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
	ginRoute.Run(util.HttpPort)
}
