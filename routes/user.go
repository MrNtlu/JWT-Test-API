package routes

import (
	"TestAPI/controllers"
	"TestAPI/db"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.RouterGroup, jwtToken *jwt.GinJWTMiddleware, mongoDB *db.MongoDB) {
	userController := controllers.NewUserController(mongoDB)

	auth := router.Group("/auth")
	{
		auth.POST("/login", jwtToken.LoginHandler)
		auth.POST("/register", userController.Register)
		auth.POST("/logout", jwtToken.LogoutHandler)
		auth.GET("/refresh", jwtToken.RefreshHandler)
	}

	user := router.Group("/user").Use(jwtToken.MiddlewareFunc())
	{
		user.GET("/info", userController.GetLoggedInUserInformation)
	}
}
