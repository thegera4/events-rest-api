package models

import "github.com/thegera4/events-rest-api/db"

type Registration struct {
	ID      int64
	UserID  int64
	EventID int64
}

func GetAllRegistrations() ([]Registration, error) {
	query := "SELECT * FROM registrations"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	registrations := []Registration{}

	for rows.Next() {
		var r Registration
		err := rows.Scan(&r.ID, &r.UserID, &r.EventID)
		if err != nil {
			return nil, err
		}

		registrations = append(registrations, r)
	}

	return registrations, nil
}