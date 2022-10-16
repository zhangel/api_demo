package middleware

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"tip/common"
)

var (
	noneAuth   = []string{"login", "logout", "token", "swagger", "docs", "favicon.ico"}
	jwtSecret  = []byte("123")
	expireTime = 3600 //second
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func checkNoneAuthURL(currURL string) bool {
	for _, noAuth := range noneAuth {
		if strings.Contains(currURL, noAuth) {
			return true
		}
	}
	return false
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims == nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GenerateToken(userId string) (string, error) {
	currentTime := time.Now()
	expireTime := currentTime.Add(time.Duration(expireTime) * time.Second)
	claims := Claims{userId, jwt.StandardClaims{ExpiresAt: expireTime.Unix(), Issuer: "gin-user"}}
	//generate token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if checkNoneAuthURL(c.Request.RequestURI) {
			return
		}
		token := c.Query("token")
		if token == "" {
			common.Json(http.StatusUnauthorized, "没有携带token", false, c)
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			common.Json(http.StatusUnauthorized, "token验证失败", false, c)
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			common.Json(http.StatusUnauthorized, "token已过期", false, c)
			c.Abort()
			return
		}
	}
}
