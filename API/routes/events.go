package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"template.com/api/models"
)

func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "could not fetch events"})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	//fetch single event
	eventId := context.Param("eventId")
	event, errs := models.GetEvent(eventId)
	if errs != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"unable to fetch event:": errs.Error()})
		return
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {
	//user authentication
	userType := context.GetString("userType")
	userId := context.GetString("userId")

	if userType != "admin" {
		context.AbortWithError(http.StatusUnauthorized, errors.New("unathorized user"))
		return
	}

	//event creation
	var newEvent models.Event
	err := context.ShouldBindJSON(&newEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event"})
		return
	}

	eventId := uuid.NewString()
	newEvent.ID = eventId
	newEvent.UserID = userId
	errs := newEvent.Save()

	if errs != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event", "error": errs.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "Event:": newEvent})
}

func updateEvent(context *gin.Context) {
	var newEvent = models.Event{}

	err := context.ShouldBindJSON(&newEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event"})
		return
	}
	errs := newEvent.UpdateEvent()

	if errs != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error:": errs.Error()})
	}

	context.JSON(http.StatusAccepted, gin.H{"Updated:": newEvent})
}

func deleteEvent(context *gin.Context) {

	eventId := context.Param("eventId")
	fetchedEvent, errs := models.GetEvent(eventId)

	if errs != nil {
		context.JSON(http.StatusBadGateway, gin.H{"unable to fetch event:": errs.Error()})
		return
	}

	userId := context.GetString("userId")

	if userId != fetchedEvent.UserID {
		context.JSON(http.StatusUnauthorized, "access denied")
		return
	}

	err := models.DeleteEvent(eventId)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error:": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"Deleted:": eventId})
}
