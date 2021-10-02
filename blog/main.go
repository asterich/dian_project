package main

import (
	"blog/model"
	"blog/routers"
	"blog/utils"
)

func main() {

	model.InitDb()

	var r = routers.InitRouter()
	r.Run(utils.HttpPort)
}
