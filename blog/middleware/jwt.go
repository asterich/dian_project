package middleware

import (
	"blog/utils"
	"blog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, errmsg.ErrCode) {
	var ExpireTime = time.Now().Add(time.Minute * time.Duration(utils.MaxLoginTime))
	var claims = MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: ExpireTime.Unix(),
			Issuer:    "blog-qsh",
		},
	}
	var reqClaim = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var token, err = reqClaim.SignedString(utils.JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCEED
}

func ParseToken(token string) (*MyClaims, errmsg.ErrCode) {
	var reqtoken, _ = jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) { return JwtKey, nil })
	var key, jwtcode = reqtoken.Claims.(*MyClaims)
	if !jwtcode || !reqtoken.Valid {
		return nil, errmsg.ERROR
	}
	return key, errmsg.SUCCEED
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		c.Set(key.Username, "username")
		c.Next()
	}
}
