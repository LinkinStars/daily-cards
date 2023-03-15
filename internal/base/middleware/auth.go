package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LinkinStars/dc/internal/base/config"
	"github.com/LinkinStars/dc/internal/base/handler"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(must bool) gin.HandlerFunc {
	unAuthHandle := func(ctx *gin.Context) {
		if must {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, handler.ErrRespBody("Unauthorized"))
		}
	}

	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			unAuthHandle(c)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.GlobalConf.Secret), nil
		})
		if err != nil {
			unAuthHandle(c)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
			c.Next()
		} else {
			unAuthHandle(c)
			return
		}
	}
}

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GlobalConf.Secret))
}
