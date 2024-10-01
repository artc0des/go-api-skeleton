package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"template.com/api/models"
	"template.com/api/utils"
)

func signup(context *gin.Context) {
	newUser := models.User{}

	err := context.ShouldBindJSON(&newUser)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	userId := uuid.NewString()
	newUser.ID = userId
	err = newUser.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"User Registered:": newUser})

}

func login(context *gin.Context) {
	user := models.User{}
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"failed to bind json": err.Error()})
		return
	}

	validUser, userId, userType, userErr := user.Validate()

	if userErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"failed to validate": userErr.Error()})
		return
	}

	if !validUser {
		context.JSON(http.StatusInternalServerError, "Invalid User")
		return
	}

	jwt, tErr := utils.GenerateToken(user.Email, userId, userType)

	if tErr != nil {
		context.JSON(http.StatusInternalServerError, tErr.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{"Authenticated": jwt})
}

func getAllUsers(context *gin.Context) {
	users, err := models.GettAllUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)

}
