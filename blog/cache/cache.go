package cache

import (
	"blog/utils"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var WhiteList *redis.Client

func StartRedis() {
	var ctx = context.TODO()
	WhiteList = redis.NewClient(&redis.Options{
		Addr:     utils.RedisAddr,
		Password: utils.RedisPassword,
		DB:       int(utils.RedisDB),
	})
	var _, err = WhiteList.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

}

func EndRedis() {
	var ctx = context.TODO()
	var res, err = WhiteList.Del(ctx, "whitelist").Result()
	if err != nil {
		fmt.Println(res, err.Error())
	}
}
