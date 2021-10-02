package main

import (
	"blog/routers"
	"blog/utils"
)

func main() {
	var r = routers.InitRouter()
	r.Run(utils.HttpPort)
}
