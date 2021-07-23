package controller

import (
	"github.com/GDGVIT/devjams21-backend/api/views"
	"github.com/GDGVIT/devjams21-backend/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthController(c *gin.Context) {
	dbHealth := db.ConnectionHealth()
	if dbHealth != nil {
		views.ErrorView(dbHealth, c)
	} else {
		views.DataView(c, http.StatusOK, "success", gin.H{
			"healthy": true,
		})
	}
}
