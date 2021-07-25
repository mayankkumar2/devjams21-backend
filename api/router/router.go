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

	teamRouter := r.Group("/team")
	{
		teamRouter.GET("/fetch", middleware.AuthMiddleware(), middleware.AttachUser, controller.GetTeamController)
		teamRouter.POST("/create", middleware.AuthMiddleware(), middleware.AttachUser, controller.CreateTeamController)
		teamRouter.PATCH("/joinCode/update", middleware.AuthMiddleware(), middleware.AttachUser, controller.UpdateTeamCodeController)
		teamRouter.POST("/join", middleware.AuthMiddleware(), middleware.AttachUser, controller.JoinTeamController)
		teamRouter.DELETE("/leave", middleware.AuthMiddleware(), middleware.AttachUser, controller.LeaveTeamController)
		teamRouter.DELETE("/member/remove", middleware.AuthMiddleware(), middleware.AttachUser, controller.RemoveMemberController)
		teamRouter.DELETE("/member/accept", middleware.AuthMiddleware(), middleware.AttachUser, controller.AcceptMemberRequestController)
	}

	r.GET("/", controller.HealthController)
	r.GET("/health", controller.HealthController)
}

func RegisterAdminRoutes(r *gin.RouterGroup) {

}
