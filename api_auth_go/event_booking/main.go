package main

import (
	"fmt"
	"net/http"
	"sayanmondal31/event_booking/db"
	"sayanmondal31/event_booking/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/event", createEvent)

	server.Run(":8080") // listening in some domain or for dev it is localhost:8080

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	fmt.Println(context, "context")
	err := context.ShouldBindJSON(&event) // kind of like scan from fmt

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data", "err": err})
		return
	}

	event.ID = 1
	event.UserId = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

}
