package cache

import (
	"context"
	"fmt"
	"os"
	"tfserver/config"

	"github.com/go-redis/redis/v8"
)

var CDb *Cache

type Cache struct {
	Client  *redis.Client
	Context context.Context
}

//初始化缓存
func InitCache() {
	CDb = new(Cache)
	CDb.Context = context.Background()
	CDb.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s%s", config.CHost, config.CPort),
		Password: config.CPassword,
		DB:       0,
	})

	_, err := CDb.Client.Ping(CDb.Context).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
		os.Exit(1)
	}
}
