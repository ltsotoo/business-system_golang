package middleware

import (
	"business-system_golang/config"
	"business-system_golang/utils/msg"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(config.SystemConfig.Server.JwtKey)

type MyClaims struct {
	EmployeeUID        string
	Permission_uid_set []string
	jwt.StandardClaims
}

//生成token
func SetToken(EmployeeUID string, Permission_uid_set []string) (token string, code int) {
	expireTime := time.Now().Add(12 * time.Hour)
	SetClaims := MyClaims{
		EmployeeUID,
		Permission_uid_set,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "zyhkAdmin",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", msg.ERROR
	}
	return token, msg.SUCCESS
}

//验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, msg.SUCCESS
	} else {
		return nil, msg.ERROR
	}
}

//api控制
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		if tokenHerder == "" {
			msg.Message(c, msg.ERRPR_TOKEN_NOT_EXIST, nil)
			c.Abort()
			return
		}
		checkToken := strings.Split(tokenHerder, " ")
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			msg.Message(c, msg.ERRPR_TOKEN_WRONG, nil)
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == msg.ERROR {
			msg.Message(c, msg.ERRPR_TOKEN_WRONG, nil)
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			msg.Message(c, msg.ERROR_TOKEN_EXPIRED, nil)
			c.Abort()
			return
		}
		c.Set("employeeUID", key.EmployeeUID)
		c.Set("Permission_uid_set", key.Permission_uid_set)
		c.Next()
	}
}
