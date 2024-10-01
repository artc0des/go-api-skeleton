package models

import (
	"template.com/api/database"
	"template.com/api/utils"
)

type User struct {
	ID       string
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Type     string
}

func (user User) Save() error {
	query := "INSERT INTO users(id, email, password, type) VALUES (?, ?, ?, ?)"
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, hErrs := utils.HashPassword(user.Password)

	if hErrs != nil {
		return hErrs
	}

	stmt.Exec(user.ID, user.Email, hashedPass, user.Type)

	return nil
}

func GettAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	users := []User{}
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Type)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (user User) Validate() (bool, string, string, error) {
	query := "SELECT id, password, type FROM users WHERE email = ?"
	row := database.DB.QueryRow(query, user.Email)

	var fetchedPass string
	var userId string
	var userType string
	err := row.Scan(&userId, &fetchedPass, &userType)

	if err != nil {
		return false, "", "", err
	}

	return utils.CheckPassHash(user.Password, fetchedPass), userId, userType, nil

}
