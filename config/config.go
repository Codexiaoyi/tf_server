package config

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

	LogConfig logConfig

	OssEndpoint        string
	OssAccessKeyId     string
	OssAccessKeySecret string
	OssBucketName      string
)

type logConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

func init() {
	file, err := ini.Load("config/config_debug.ini")
	if err != nil {
		fmt.Println("Load config file error!", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadLog(file)
	LoadOSS(file)
}

//加载服务器设置
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	jwtKeyStr := file.Section("server").Key("JwtKey").MustString("asdasccuiwej8498as5123xzc")
	JwtKey = []byte(jwtKeyStr)
	JwtIssuer = file.Section("server").Key("JwtIssuer").MustString("travelfriend")
}

//加载数据库配置
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("tfserver")
	DbPassword = file.Section("database").Key("DbPassword").MustString("tfserver")
	DbName = file.Section("database").Key("DbName").MustString("tfserver")
}

//加载日志配置
func LoadLog(file *ini.File) {
	LogConfig.Level = file.Section("log").Key("Level").MustString("debug")
	LogConfig.Filename = file.Section("log").Key("Filename").MustString("tf.log")
	LogConfig.MaxSize = file.Section("log").Key("MaxSize").MustInt(200)
	LogConfig.MaxAge = file.Section("log").Key("MaxAge").MustInt(7)
	LogConfig.MaxBackups = file.Section("log").Key("MaxBackups").MustInt(10)
}

//加载oss相关配置
func LoadOSS(file *ini.File) {
	OssEndpoint = file.Section("aliOss").Key("Endpoint").MustString("<yourEndpoint>")
	OssAccessKeyId = file.Section("aliOss").Key("AccessKeyId").MustString("<yourAccessKeyId>")
	OssAccessKeySecret = file.Section("aliOss").Key("AccessKeySecret").MustString("<yourAccessKeySecret>")
	OssBucketName = file.Section("aliOss").Key("BucketName").MustString("<yourBucketName>")
}
