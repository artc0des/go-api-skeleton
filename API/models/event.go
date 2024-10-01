package models

import (
	"time"

	"template.com/api/database"
)

type Event struct {
	ID          string
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      string
}

var cache = []Event{}

func (event Event) Save() error {
	//persist in DB
	query := `
	INSERT INTO events(id, name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?, ?)`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, error := stmt.Exec(event.ID, event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	//events = append(events, event)
	if error != nil {
		return err
	}

	cache = append(cache, event)
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	var events = []Event{}
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEvent(eventId string) (Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	row := database.DB.QueryRow(query, eventId)

	readEvent := Event{}

	err := row.Scan(&readEvent.ID, &readEvent.Name, &readEvent.Description, &readEvent.Location, &readEvent.DateTime, &readEvent.UserID)

	if err != nil {
		return Event{}, err
	}

	return readEvent, nil
}

func (event *Event) UpdateEvent() error {
	query := `UPDATE events
	SET name = ?
	WHERE id = ?`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, error := stmt.Exec(event.Name, event.ID)

	if error != nil {
		return err
	}

	cache = append(cache, *event)
	return nil
}

func DeleteEvent(eventId string) error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(eventId)

	if err != nil {
		return err
	}

	return nil
}

func (event Event) Register(userId, registrationId string) error {
	query := "INSERT INTO registrations(id, event_id, user_id) VALUES (?,?,?)"
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(registrationId, event.ID, userId)

	if err != nil {
		return err
	}

	return nil
}

func CancelRegistration(userId, eventId string) error {
	query := "DELETE FROM registrations WHERE user_id = ? AND event_id = ?"

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, eventId)

	if err != nil {
		return err
	}

	return nil
}
