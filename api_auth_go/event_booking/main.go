package main

import (
	"net/http"
	"sayanmondal31/event_booking/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // listening in some domain or for dev it is localhost:8080

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	context.ShouldBindJSON(&event)

}
