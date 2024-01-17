package service

import (
	"database/sql"
	"echo-api/db"
	"log"
	"net/http"
)

type Addresses struct {
	Id               int    `json:"id"`
	User_Id          int    `json:"user_id"`
	User_Name        string `json:"user_name"`
	Street_Addresses string `json:"street_addresses"`
	City             string `json:"city"`
	Province         string `json:"province"`
	Country          string `json:"country"`
	Phone_Number     string `json:"phone_number"`
}

func GetAddressesUsers(User_Name string) (Response, error) {
	var add Addresses
	var res Response
	cond := db.NewData()
	query := "SELECT * FROM addresses WHERE user_name = ?"
	rows, err := cond.Query(query, User_Name)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&add.Id, &add.User_Id, &add.User_Name, &add.Street_Addresses,
			&add.City, &add.Province, &add.Country, &add.Phone_Number)

		if err == sql.ErrNoRows {
			return res, sql.ErrNoRows
		}
		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = add
	}

	return res, nil
}

func CreateAddresses(User_Id int, User_Name, Street_Addresses, City, Province, Country, Phone_Number string) (Response, error) {
	var res Response
	cond := db.NewData()
	query := "INSERT INTO addresses (user_id,user_name,street_addresses,city,province,country,phone_number) VALUES (?,?,?,?,?,?,?)"
	sqlstmnt, err := cond.Prepare(query)
	if err != nil {
		log.Println("Query Error")
		return res, err
	}
	result, err := sqlstmnt.Exec(&User_Id, &User_Name, &Street_Addresses, &City, &Province, &Country, &Phone_Number)
	if err != nil {
		log.Println("Exec Error")
		return res, err
	}
	lastinsert, err := result.LastInsertId()
	if err != nil {
		log.Printf("No result because %s", err)
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "Success Create"
	res.Data = map[string]int64{
		"lastinsertId": lastinsert,
	}
	return res, err
}
