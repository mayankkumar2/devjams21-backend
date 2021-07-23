package router

import (
	"github.com/GDGVIT/devjams21-backend/api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.Engine)  {
	usrRouter := r.Group("/user")
	{
		usrRouter.GET("", controller.HealthController)
	}


	r.GET("/", controller.HealthController)
	r.GET("/health", controller.HealthController)
}

func RegisterAdminRoutes(r *gin.Engine)  {

}