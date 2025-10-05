package models

import (
	"sayanmondal31/event_booking/db"
	"sayanmondal31/event_booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users(email,password) VALUES (?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	haspass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, haspass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err

}

func (u *User) VerifyPass() error {
	query := `
		SELECT  id,password FROM users Where email = ?
	`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedHashPassword string
	err := row.Scan(&u.ID, &retrievedHashPassword)
	if err != nil {
		return err
	}

	err = utils.DecryptHash(retrievedHashPassword, u.Password)

	if err != nil {
		return err
	}

	return nil
}
