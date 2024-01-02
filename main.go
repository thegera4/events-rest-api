package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/events-rest-api/db"
	"github.com/thegera4/events-rest-api/models"
)

//NOTE 1 framework for rest api: go get -u github.com/gin-gonic/gin
//NOTE 2 use the sql package from go + a sqldriver (in this case for sqlite3) go get modernc.org/sqlite
//NOTE 3 run "set CGO_ENABLED=1" if you are on windows before "go run ." because sqlite3 needs this
func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events!"})
		return
	}
	context.JSON(http.StatusOK, events) //you can pass the status code as a number (200)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event) //store data from the body in "event" variable, can return err
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	event.ID = 1 //change it later
	event.UserID = 1 //change it later

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "Event created successfully!", "event": event})
}