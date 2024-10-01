package models

type Registration struct {
	ID       string
	Event_id string `binding:"required"`
	User_id  string
}
