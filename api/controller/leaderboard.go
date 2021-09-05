package controller

import (
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLeaderboard(ctx *gin.Context)  {
	e, _ := db.LeaderboardService.GetLeaderBoard(ctx)
	views.DataView(ctx, http.StatusOK, "success", e)
}
