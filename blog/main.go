package main

import (
	"blog/cache"
	"blog/model"
	"blog/routers"
	"blog/utils"
)

func main() {

	cache.StartRedis()
	model.InitDb()

	var r = routers.InitRouter()
	r.Run(utils.HttpPort)

	cache.EndRedis()
}
