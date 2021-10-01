package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func addTimestampApi(c *gin.Context) {
	var path = c.FullPath()
	fmt.Println(path)
	var jsonstr = c.DefaultQuery("json", "{}")
	var obj gin.H
	json.Unmarshal([]byte(jsonstr), &obj)
	obj["timestamp"] = time.Now().Unix()
	c.JSON(http.StatusOK, obj)
}

func setupRouter() *gin.Engine {
	var r = gin.Default()

	r.GET("/hello", addTimestampApi)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
