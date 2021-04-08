package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"tfserver/config"
	"tfserver/util/log"

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

//set对象，json序列化
func (c *Cache) SetObject(key, field string, object interface{}) error {
	data, _ := json.Marshal(object)
	err := c.Client.HSet(c.Context, key, field, data).Err()
	if err != nil {
		log.DebugLog("redis错误", err.Error())
	}
	return err
}

//get对象，json序列化
func (c *Cache) GetObject(key, field string, object interface{}) error {
	result, _ := c.Client.HGet(c.Context, key, field).Bytes()
	err := json.Unmarshal(result, object)
	if err != nil {
		log.DebugLog("redis错误", err.Error())
	}
	return err
}

//set
func (c *Cache) Set(key, field string, object interface{}) error {
	err := c.Client.HSet(c.Context, key, field, object).Err()
	if err != nil {
		log.DebugLog("redis错误", err.Error())
	}
	return err
}

//get
func (c *Cache) Get(key, field string) ([]byte, error) {
	result, err := c.Client.HGet(c.Context, key, field).Bytes()
	if err != nil {
		log.DebugLog("redis错误", err.Error())
	}
	return result, err
}

//判读缓存是否存在
func (c *Cache) IsExist(key, field string) bool {
	result, err := CDb.Client.HExists(c.Context, key, field).Result()
	if err != nil {
		log.DebugLog("redis错误", err.Error())
		return false
	}

	return result
}

//删除缓存
func (c *Cache) Delete(key, field string) bool {
	result, err := CDb.Client.HDel(c.Context, key, field).Result()
	if err != nil {
		log.DebugLog("redis错误", err.Error())
		return false
	}

	if result > 0 {
		return true
	} else {
		return false
	}
}
