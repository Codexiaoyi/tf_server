package util

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode   string
	HttpPort  string
	JwtKey    []byte
	JwtIssuer string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Load config file error!", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	jwtKeyStr := file.Section("server").Key("JwtKey").MustString("asdasccuiwej8498as5123xzc")
	JwtKey = []byte(jwtKeyStr)
	JwtIssuer = file.Section("server").Key("JwtIssuer").MustString("travelfriend")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("tfserver")
	DbPassword = file.Section("database").Key("DbPassword").MustString("tfserver")
	DbName = file.Section("database").Key("DbName").MustString("tfserver")
}
