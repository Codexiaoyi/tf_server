package main

import (
	"tfserver/infrastructure"
	"tfserver/route"
	"tfserver/util"
)

func main() {
	//数据库初始化
	infrastructure.InitDbContext()
	//初始化gin的路由
	ginRoute := route.InitRouter()
	//运行
	ginRoute.Run(util.HttpPort)
}
