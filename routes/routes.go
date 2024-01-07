package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thegera4/events-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)
	server.GET("/events/registrations", getAllEventsRegistrations)

	authenticated := server.Group("/") //create a new group for auth protected routes
	authenticated.Use(middlewares.Authenticate) //use the Authenticate middleware to protect routes
	authenticated.POST("/events", createEvent) //routes in this group are now protected
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}