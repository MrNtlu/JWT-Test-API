package controllers

import (
	"TestAPI/db"
	"TestAPI/models"
	"TestAPI/requests"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Database *db.MongoDB
}

func NewUserController(mongoDB *db.MongoDB) UserController {
	return UserController{
		Database: mongoDB,
	}
}

var (
	errAlreadyRegistered = "User already registered."
	errPasswordNoMatch   = "Passwords do not match."
	errNoUser            = "Sorry, couldn't find user."
)

func (u *UserController) Register(c *gin.Context) {
	var data requests.Register
	if shouldReturn := bindJSONData(&data, c); shouldReturn {
		return
	}

	userModel := models.NewUserModel(u.Database)
	user, _ := userModel.FindUserByEmail(data.EmailAddress)

	if user.EmailAddress != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errAlreadyRegistered,
		})

		return
	}

	if err := userModel.CreateUser(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registered successfully."})
}

func (u *UserController) GetLoggedInUserInformation(c *gin.Context) {
	uid := jwt.ExtractClaims(c)["id"].(string)

	userModel := models.NewUserModel(u.Database)
	userInfo, _ := userModel.FindUserByID(uid)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully fetched user info.", "data": userInfo})
}
