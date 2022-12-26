package main

import (
	"TestAPI/db"
	"TestAPI/helpers"
	"TestAPI/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
)

func main() {
	if os.Getenv("ENV") != "Production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Default().Println(os.Getenv("ENV"))
			log.Fatal("Error loading .env file")
		}
	}

	mongoDB, ctx, cancel := db.Connect(os.Getenv("MONGO_ATLAS_URI"))
	defer db.Close(ctx, mongoDB.Client, cancel)

	jwtHandler := helpers.SetupJWTHandler(mongoDB)

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	const (
		burstTime       = 100 * time.Millisecond
		requestCount    = 20
		restrictionTime = 5 * time.Second
	)
	// Burst of 0.1 sec 20 requests. 5 second restriction.
	router.Use(limit.NewRateLimiter(func(ctx *gin.Context) string {
		return ctx.ClientIP()
	}, func(ctx *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(burstTime), requestCount), restrictionTime
	}, func(ctx *gin.Context) {
		const tooManyRequestError = "Too many requests. Rescricted for 5 seconds."
		ctx.JSON(http.StatusTooManyRequests, gin.H{"error": tooManyRequestError, "message": tooManyRequestError})
		ctx.Abort()
	}))

	routes.SetupRoutes(router, jwtHandler, mongoDB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
