package models

import (
	"sayanmondal31/event_booking/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() error {
	// later: add it to database
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() []Event {
	return events
}
