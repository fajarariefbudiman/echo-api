package service

import (
	"database/sql"
	"echo-api/db"
	"net/http"

	"github.com/go-playground/validator"
)

type Users struct {
	Id         int64  `json:"id"`
	Name       string `validate:"required" json:"name"`
	Email      string `validate:"required|email" json:"email"`
	Password   string `validate:"required" json:"password"`
	Handphone  string `validate:"required" json:"handphone"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
	Addresses  Addresses
}

func GetAllUsers() (Response, error) {
	var get Users
	var getAll []Users
	var res Response

	cond := db.NewData()
	sqlstmt := "SELECT * FROM users"

	rows, err := cond.Query(sqlstmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Name, &get.Email, &get.Password, &get.Handphone, &get.Created_At, &get.Updated_At)
		if err != nil {
			return res, err
		}
		getAll = append(getAll, get)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = getAll

	return res, nil
}

func CreateUsers(name, email, password, handphone string) (Response, error) {
	var response Response
	v := validator.New()
	Use := Users{
		Name:      name,
		Email:     email,
		Password:  password,
		Handphone: handphone,
	}
	err := v.Struct(Use)
	if err != nil {
		return response, err
	}
	cond := db.NewData()
	sqlstmnt := "INSERT INTO users (name,email,password,handphone) VALUES(?,?,?,?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(name, email, password, handphone)
	if err != nil {
		return response, err
	}
	lastinsert, err := result.LastInsertId()
	if err != nil {
		return response, err
	}
	response.Status = http.StatusCreated
	response.Message = "Success Create"
	response.Data = map[string]int64{
		"lastinsertId": lastinsert,
	}
	return response, err
}

func DeleteUsers(Id int) (Response, error) {
	var response Response
	cond := db.NewData()
	sqlstmnt := "DELETE FROM users WHERE id=?"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}

	result, err := stmt.Exec(Id)
	if err != nil {
		return response, err
	}

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return response, err
	}

	response.Status = http.StatusNoContent
	response.Message = "Success Delete"
	response.Data = map[string]int64{
		"rowsaffected": rowsaffected,
	}
	return response, err

}

func UpdateUsers(Id int, Name, Email, Password, Handphone string) (Response, error) {
	var response Response
	cond := db.NewData()
	sqlstmnt := "UPDATE users set name=?, email=?, password=?, handphone=? WHERE id=?"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(Name, Email, Password, Handphone, Id)
	if err != nil {
		return response, err
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return response, err
	}
	response.Status = http.StatusOK
	response.Message = "Succes Update"
	response.Data = map[string]int64{
		"rowsaffected": rowsaffected,
	}
	return response, err
}

func GetUsersId(Id int) (Response, error) {
	var get Users
	var res Response
	cond := db.NewData()
	sqlstmt := "SELECT id,name,email,password,handphone,created_at,updated_at FROM users WHERE id=?"
	rows, err := cond.Query(sqlstmt, Id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Name, &get.Email, &get.Password, &get.Handphone, &get.Created_At, &get.Updated_At)
		if err == sql.ErrNoRows {
			return res, sql.ErrNoRows
		}
		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = get

	}
	return res, nil
}
