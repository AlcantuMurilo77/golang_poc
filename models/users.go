package models

import (
	"example.com/my_go_proj/db"
	"example.com/my_go_proj/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
  INSERT INTO users(email, password)
  VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPass

	result, err := stmt.Exec(u.Email, hashedPass)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id

	return err
}
