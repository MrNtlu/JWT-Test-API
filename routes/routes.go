package routes

import (
	"TestAPI/db"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, jwtToken *jwt.GinJWTMiddleware, mongoDB *db.MongoDB) {
	apiRouter := router.Group("/api")
	userRouter(apiRouter, jwtToken, mongoDB)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "All routes lead to rome"})
	})
}
