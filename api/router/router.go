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
		usrRouter.GET("/participation", middleware.AuthMiddleware(), middleware.AttachUser, controller.UserParticipationController)
	}

	usrInfoRouter := r.Group("/user_info")
	{
		usrInfoRouter.POST("/create", middleware.AuthMiddleware(), middleware.AttachUser, controller.CreateUserInfoController)
		usrInfoRouter.PATCH("/update", middleware.AuthMiddleware(), middleware.AttachUser, controller.UpdateUserInfoController)
		usrInfoRouter.DELETE("/delete", middleware.AuthMiddleware(), middleware.AttachUser, controller.DeleteUserInfoController)
		usrInfoRouter.GET("/get", middleware.AuthMiddleware(), middleware.AttachUser, controller.GetUserInfoController)
	}

	teamRouter := r.Group("/team")
	{
		teamRouter.GET("/fetch", middleware.AuthMiddleware(), middleware.AttachUser, controller.GetTeamController)
		teamRouter.POST("/create", middleware.AuthMiddleware(), middleware.AttachUser, controller.CreateTeamController)
		teamRouter.PATCH("/joinCode/update", middleware.AuthMiddleware(), middleware.AttachUser, controller.UpdateTeamCodeController)
		teamRouter.PATCH("/teamName", middleware.AuthMiddleware(), middleware.AttachUser, controller.UpdateTeamNameController)
		teamRouter.POST("/join", middleware.AuthMiddleware(), middleware.AttachUser, controller.JoinTeamController)
		teamRouter.DELETE("/leave", middleware.AuthMiddleware(), middleware.AttachUser, controller.LeaveTeamController)
		teamRouter.DELETE("/member/remove", middleware.AuthMiddleware(), middleware.AttachUser, controller.RemoveMemberController)
		teamRouter.PATCH("/member/accept", middleware.AuthMiddleware(), middleware.AttachUser, controller.AcceptMemberRequestController)
	}

	participationRouter := r.Group("/participation")
	{
		participationRouter.DELETE("/remove", middleware.AuthMiddleware(), middleware.AttachUser, controller.DeleteParticipationController)
		participationRouter.GET("/teams/:event_id", controller.GetTeamsController)
		participationRouter.POST("/create", middleware.AuthMiddleware(), middleware.AttachUser, controller.CreateParticipationController)
	}
	eventRouter := r.Group("/event")
	{
		eventRouter.GET("/all", controller.GetAllEventsController)
		eventRouter.GET("/fetch/:event_id", middleware.AuthMiddleware(), middleware.AttachUser, controller.GetEventController)
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
		eventRouter.GET("/:event_id", controller.GetEventController)
	}

	challengeRouter := r.Group("/challenge")
	{
		challengeRouter.POST("/create", controller.CreateChallengeController)
		challengeRouter.GET("/get", controller.GetChallengeController)
		challengeRouter.PUT("/update", controller.UpdateChallengeController)
		challengeRouter.DELETE("/delete", controller.DeleteChallengeController)
	}

}
