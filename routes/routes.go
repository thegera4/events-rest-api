package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thegera4/events-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)

	authenticated := server.Group("/") //create a new group for auth protected routes
	authenticated.Use(middlewares.Authenticate) //use the Authenticate middleware to protect routes
	authenticated.POST("/events", createEvent) //routes in this group are now protected
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}