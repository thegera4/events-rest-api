package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)

	server.POST("/events", createEvent)
	
	server.PUT("/events/:id", updateEvent)
}