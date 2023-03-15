package clint

import (
	"encoding/json"
	"github.com/go-redis/redis"
	config2 "go-blog/config"
	"time"
)

var Rd *redis.Client

func init() {
	var configObj config2.RedisConfig
	configInfo := configObj.GetRedisConfig()

	Rd = redis.NewClient(&redis.Options{
		Addr:         configInfo.Host + ":" + configInfo.Port,
		Password:     configInfo.Pass,
		DB:           configInfo.DbNumber,
		ReadTimeout:  time.Millisecond * time.Duration(100),
		WriteTimeout: time.Millisecond * time.Duration(100),
		IdleTimeout:  time.Second * time.Duration(60),
	})
	_, err := Rd.Ping().Result()
	if err != nil {
		panic("init redis error")
	}
}

//缓存获取
func GetCache(key string) (string, bool) {
	r, err := Rd.Get(key).Result()
	if err != nil {
		return "", false
	}
	return r, true
}

//缓存设置
func SetCache(key string, val string, expTime int32) {
	Rd.Set(key, val, time.Duration(expTime)*time.Second)
}

// 获取json缓存 然后json转map函数，通用
func JsonToMap(key string) map[string]interface{} {
	str, status := GetCache(key)
	if status != true {
		panic(status)
	}
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}
