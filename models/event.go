package models

import (
	"time"
	"github.com/thegera4/events-rest-api/db"
)

type Event struct {
	ID          int64   
	Title       string 	  `binding:"required"`
	Description string 	  `binding:"required"`
	Location    string 	  `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int   
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events (title, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query) //Prepare can lead to better performance instead of Exec the query directly
	if err != nil {
		return err
	}
	defer stmt.Close() //close the statement after the function ends

	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID) //Exec to update stuff
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() //to get the last inserted id
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events" //since it is a simple query, you do not need to prepare it
	rows, err := db.DB.Query(query) //Query is used to get/fetch a bunch of data/rows
	if err != nil {
		return nil, err
	}
	defer rows.Close() //close the rows after the function ends

	events := []Event{}

	for rows.Next() { //loop through the rows
		var e Event
		err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.DateTime, &e.UserID) //scan the rows and store it in the variable
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id) //QueryRow is used to get/fetch a single row

	var e Event
	err := row.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET title = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Title, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}