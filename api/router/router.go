package router

import (
	"github.com/GDGVIT/devjams21-backend/api/controller"
	"github.com/GDGVIT/devjams21-backend/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup) {
	usrRouter := r.Group("/user")
	{
		usrRouter.POST("/create", controller.CreateUserController)
		usrRouter.POST("/login", controller.UserLoginController)
		usrRouter.GET("/profile", middleware.AuthMiddleware(), middleware.AttachUser, controller.UserProfileController)
		usrRouter.PATCH("/update", middleware.AuthMiddleware(), middleware.AttachUser, controller.UserProfileUpdateController)
		usrRouter.GET("/teams", middleware.AuthMiddleware(), middleware.AttachUser, controller.UserTeamsController)
		usrRouter.GET("/leader", middleware.AuthMiddleware(), middleware.AttachUser, controller.UserLeaderController)
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

	partRouter := r.Group("/part")
	{
		partRouter.GET("/teams", middleware.AuthMiddleware(), middleware.AttachUser, controller.GetTeamsController)
	}

	r.GET("/", controller.HealthController)
	r.GET("/health", controller.HealthController)
}

func RegisterAdminRoutes(r *gin.RouterGroup) {

	eventRouter := r.Group("/event")
	{
		eventRouter.POST("/create", controller.CreateEventController)
		eventRouter.PUT("/update", controller.UpdateEventController)
		eventRouter.DELETE("/delete", controller.DeleteEventController)
		eventRouter.GET("/get", controller.GetEventController)
	}

	challengeRouter := r.Group("/challenge")
	{
		challengeRouter.POST("/create", controller.CreateChallengeController)
		challengeRouter.GET("/get", controller.GetChallengeController)
		challengeRouter.PUT("/update", controller.UpdateChallengeController)
		challengeRouter.DELETE("/delete", controller.DeleteChallengeController)
	}

}
