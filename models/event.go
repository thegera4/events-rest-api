package models

import (
	"github.com/thegera4/events-rest-api/db"
)

type Event struct {
	ID          int64   
	Title       string 	  `binding:"required"`
	Description string 	  `binding:"required"`
	Location    string 	  `binding:"required"`
	Date    	string 	  `binding:"required"`
	ImageURL	string
	UserID      int64   
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (title, description, location, date_time, image_url, user_id)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query) //Prepare can lead to better performance instead of Exec the query directly
	if err != nil {
		return err
	}
	defer stmt.Close() //close the statement after the function ends

	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.Date, e.ImageURL, e.UserID) //Exec to update stuff
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
		err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.Date, &e.ImageURL, &e.UserID) //scan the rows and store it in the variable
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func FilteredEvents(year *string, month *string) ([]Event, error) {
	query := "SELECT * FROM events WHERE date_time LIKE ?"
	rows, err := db.DB.Query(query, *year+"-"+*month+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.Date, &e.ImageURL, &e.UserID)
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
	err := row.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.Date, &e.ImageURL, &e.UserID)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET title = ?, description = ?, location = ?, date_time = ?, image_url = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Title, event.Description, event.Location, event.Date, event.ImageURL, event.ID)
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

func GetEventsForUser(userId int64) ([]Event, error) {
	query := "SELECT * FROM registrations WHERE user_id = ?"
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.UserID)
		if err != nil {
			return nil, err
		}

		event, err := GetEventById(e.ID)
		if err != nil {
			return nil, err
		}

		events = append(events, *event)
	}

	return events, nil
}

func (e Event) Register(userId int64) error {
	events, err := GetEventsForUser(userId)
	if err != nil {
		return err
	}

	for _, event := range events {
		if event.ID == e.ID {
			return nil
		}
	}

	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error { 
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}