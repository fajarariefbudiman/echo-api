package service

import (
	"echo-api/db"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `validate:"required|email" json:"email"`
	Password   string `validate:"required" json:"password"`
	Handphone  string `json:"handphone"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}

func AuthenticateUser(email, password string) (bool, error) {
	// Mengambil user dari database berdasarkan email
	v := validator.New()
	Use := User{
		Email:    email,
		Password: password,
	}
	err := v.Struct(Use)
	if err != nil {
		return false, err
	}
	var user User
	con := db.NewData()
	sqlstmt := "SELECT * FROM users WHERE email=?"
	err = con.QueryRow(sqlstmt, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Handphone, &user.Created_At, &user.Updated_At)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpdatePassword(email, newPassword string) error {
	db := db.NewData()
	query := "UPDATE users SET password = ? WHERE email = ?"
	_, err := db.Exec(query, newPassword, email)
	if err != nil {
		return err
	}

	return nil
}
