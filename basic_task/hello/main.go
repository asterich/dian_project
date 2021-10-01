package main

import (
	//	"encoding/json"
	//	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

var json gin.H

func sendJsonApi(c *gin.Context) {
	json = gin.H{}
	if err := c.BindJSON(&json); err != nil {
		log.Fatal(err.Error())
	}
}

func addTimestampApi(c *gin.Context) {
	json["timestamp"] = time.Now().Unix()
	c.JSON(http.StatusOK, json)
}

func setupRouter() *gin.Engine {
	var r = gin.Default()

	r.POST("/sendjson", sendJsonApi)
	r.GET("/hello", addTimestampApi)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
