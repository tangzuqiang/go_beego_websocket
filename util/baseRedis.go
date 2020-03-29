package utils

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"strconv"
)

var BaseRedis baseRedisModel

type baseRedisModel struct{}

func (baseRedisModel) Connect(Key string) *redis.Client {
	db, _ := strconv.Atoi(beego.AppConfig.String("redis_db"))
	redisClient := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis_addr"),
		Password: beego.AppConfig.String("redis_password"),
		DB:       db,
	})

	return redisClient
}

func (baseRedisModel) Close(redisClient *redis.Client) {
	redisClient.Close()
}
