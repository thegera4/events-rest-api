package routes

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/events-rest-api/models"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event!"})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event or You are already registered!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered successfully!"})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled successfully!"})
}

func getAllEventsRegistrations(context *gin.Context) {
	events, err := models.GetAllRegistrations()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch registrations!"})
		return
	}

	context.JSON(http.StatusOK, events)
}