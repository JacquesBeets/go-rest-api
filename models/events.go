package models

import (
	"time"

	"github.com/jacquesbeets/go-rest-api/db"
)

type Event struct {
	ID          int64       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"userId"`
}



func (e Event) Save() error {
	query := `
	INSERT INTO events (title, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	// Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).
	// This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	e.ID = id
	return err
}

func GetEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
