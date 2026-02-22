package main

import (
	"example.com/my_go_proj/db"
	"example.com/my_go_proj/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080")
	if err != nil {
		panic("Could not start server")
	}
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse req data"})
		return
	}

	event.ID = 1     //FOr now a dummy value
	event.UserID = 1 //here aswell

	models.Event.Save(event)
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
