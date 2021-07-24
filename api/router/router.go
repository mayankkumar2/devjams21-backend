package router

import (
	"github.com/GDGVIT/devjams21-backend/api/controller"
	"github.com/GDGVIT/devjams21-backend/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup) {
	usrRouter := r.Group("/user")
	{
		usrRouter.GET("/", controller.HealthController)
		usrRouter.POST("/create", controller.CreateUserController)
		usrRouter.POST("/login", controller.UserLoginController)
		usrRouter.GET("/profile", middleware.AuthMiddleware(), middleware.AttachUser, controller.UserProfileController)
		usrRouter.PATCH("/update", middleware.AuthMiddleware(), middleware.AttachUser, controller.UserProfileUpdateController)
	}

	r.GET("/", controller.HealthController)
	r.GET("/health", controller.HealthController)
}

func RegisterAdminRoutes(r *gin.RouterGroup) {

}
