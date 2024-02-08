package models

import (
	"errors"

	"github.com/jacquesbeets/go-rest-api/db"
	"github.com/jacquesbeets/go-rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `
    INSERT INTO users (email, password) 
    VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()

	u.ID = userID
	return err
}

func (u User) FindByEmailAndPassword() error {
	query := `
	SELECT id, password
	FROM users
	WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	paswordIsValid := utils.VerifyPassword(retrievedPassword, u.Password)

	if !paswordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
