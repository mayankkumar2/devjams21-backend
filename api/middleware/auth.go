package middleware

import (
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"net/http"
)

func Validate(h http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt_secret")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return jwtMiddleware.Handler(h)
}

func ValidateAndGetClaims(ctx context.Context, role string) (map[string]interface{}, error) {
	_, _ = ctx.Value("user").(*jwt.Token)

	return nil, nil
}
