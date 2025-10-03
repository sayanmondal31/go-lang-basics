package models

import (
	"time"
)

type Event struct {
	ID          int
	Name        string    `json: name, binding:"required"`
	Description string    `json: description, binding:"required"`
	Location    string    `json: location, binding:"required"`
	DateTime    time.Time `json: dateTime, binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	// later: add it to database
	events = append(events, e)

}

func GetAllEvents() []Event {
	return events
}
