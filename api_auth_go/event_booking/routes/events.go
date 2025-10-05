package routes

import (
	"fmt"
	"net/http"
	"sayanmondal31/event_booking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createEvent(context *gin.Context) {

	var event models.Event
	fmt.Println(context, "context")
	err := context.ShouldBindJSON(&event) // kind of like scan from fmt

	fmt.Println(event, "event data check")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data", "err": err})
		return
	}

	event.UserId = context.GetInt64("userId")

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	fmt.Println(err, "err")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. try again later"})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event "})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event "})
		return
	}

	context.JSON(http.StatusOK, event)

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event "})
		return
	}

	userId := context.GetInt64("userId")
	fmt.Println(userId, "user id check")
	event, err := models.GetEventById(eventId)

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update event!"})
		return

	}

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get event "})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data", "err": err})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusAccepted, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event "})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get event "})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete event!"})
		return

	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get event "})
		return
	}

	var deletedEvent models.Event

	deletedEvent.ID = eventId

	err = deletedEvent.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete event", "err": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event delete successfully"})
}
