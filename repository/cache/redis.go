package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"tfserver/config"
	"tfserver/util/log"
	"time"

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
		log.ErrorLog("连接redis出错..", err.Error())
		os.Exit(1)
	}
}

//set对象
func (c *Cache) SetObject(key string, object interface{}, expiration time.Duration) {
	data, _ := json.Marshal(object)
	err := c.Client.Set(c.Context, key, string(data), expiration).Err()
	if err != nil {
		log.DebugLog("redis错误", err.Error())
	}
}

//get对象
func (c *Cache) GetObject(key string, object interface{}) {
	result, _ := c.Client.Get(c.Context, key).Result()
	err := json.Unmarshal([]byte(result), object)
	if err != nil {
		log.DebugLog("redis错误", err.Error())
	}
}
