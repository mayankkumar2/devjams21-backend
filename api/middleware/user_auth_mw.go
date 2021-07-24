package middleware

import (
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AttachUser(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	id, _ := uuid.Parse(claims["id"].(string))
	usr, err := db.UserService.FindByID(ctx, &id)
	if err != nil {
		views.ErrorView(err, ctx)
		return
	}
	ctx.Set("user", usr)
	ctx.Next()
}
