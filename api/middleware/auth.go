package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const identityKey = "id"
var jwtMW *jwt.GinJWTMiddleware

func AuthMiddleware() gin.HandlerFunc {
	if jwtMW == nil {
		secret := os.Getenv("SECRET")
		var err error
		jwtMW, err = jwt.New(&jwt.GinJWTMiddleware{
			Key:           []byte(secret),
			Timeout:       time.Hour * 24 * 7,
			MaxRefresh:    time.Hour * 24 * 7,
			IdentityKey:   identityKey,
			TokenLookup:   "header: Authorization, query: token",
			TokenHeadName: "Bearer",
			TimeFunc:      time.Now,
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				m := data.(map[string]string)
				return jwt.MapClaims{
					"id": m["id"],
				}
			},
		})
		if err != nil {
			logrus.Errorln(err)
		}
	}
	return jwtMW.MiddlewareFunc()
}

func Token(id *uuid.UUID) (string, time.Time, error) {
	return jwtMW.TokenGenerator(map[string] string {
		"id": id.String(),
	})
}