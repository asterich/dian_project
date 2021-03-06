package middleware

import (
	"blog/cache"
	"blog/utils"
	"blog/utils/errmsg"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, errmsg.ErrCode) {
	var duration = time.Duration(utils.MaxLoginTime) * time.Minute
	log.Println(utils.MaxLoginTime, " ", duration)
	var ExpireTime = time.Now().Add(duration)
	var claims = MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: ExpireTime.Unix(),
			Issuer:    "blog-qsh",
		},
	}
	var reqClaim = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var token, err = reqClaim.SignedString([]byte(utils.JwtKey))
	if err != nil {
		log.Println(err)
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCEED
}

func ParseToken(token string) (*MyClaims, errmsg.ErrCode) {
	var reqtoken, err = jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) { return JwtKey, nil })
	if key, _ := reqtoken.Claims.(*MyClaims); reqtoken.Valid {
		return key, errmsg.SUCCEED
	}
	log.Println(err.Error())
	return nil, errmsg.ERROR
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = context.TODO()
		var tokenHeader = c.Request.Header.Get("Authorization")
		var code errmsg.ErrCode
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		var chkToken = strings.SplitN(tokenHeader, " ", 2)
		if len(chkToken) != 2 || chkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		var key, chkcode = ParseToken(chkToken[1])
		if chkcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_OUT_OF_DATE
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if res, err := cache.WhiteList.Get(ctx, fmt.Sprintf("whitelist_%s", key.Username)).Result(); res == string(redis.Nil) || err != nil {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set(key.Username, "username")
		c.Next()
	}
}
