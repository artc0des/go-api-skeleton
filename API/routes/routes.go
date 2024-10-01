package routes

import (
	"github.com/gin-gonic/gin"
	"template.com/api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getAllEvents)
	server.GET("/events/:eventId", getEvent)

	//logged in users only
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events", updateEvent)
	authenticated.DELETE("/events/:eventId", deleteEvent)
	authenticated.POST("/events/:eventId/register", registerForEvent)
	authenticated.DELETE("/events/:eventId/register", cancelRegistration)

	server.POST("/signup", signup)
	server.GET("/", login)
	server.GET("/users", getAllUsers)
	server.POST("/login", login)
}
