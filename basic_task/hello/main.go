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
	json["timestamp"] = time.Now().Unix()
	c.JSON(http.StatusOK, gin.H{
		"data": json,
	})
}

func setupRouter() *gin.Engine {
	var r = gin.Default()

	r.POST("/hello", sendJsonApi)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
