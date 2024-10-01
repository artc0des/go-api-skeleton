package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"template.com/api/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetString("userId")
	eventId := context.Param("eventId")

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, "failed to fetch")
	}

	registrationId := uuid.NewString()
	err = event.Register(userId, registrationId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, "failed to register")
		return
	}

	context.JSON(http.StatusOK, "you have been registered")
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetString("userId")
	eventId := context.Param("eventId")
	err := models.CancelRegistration(userId, eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, "failed to register")
		return
	}

	context.JSON(http.StatusOK, "event canceled")
}
