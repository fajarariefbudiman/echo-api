package service

import (
	"echo-api/db"
	"fmt"
	"log"
	"net/http"
)

type Home struct {
	TableName string `json:"name"`
	URL       string `json:"url"`
}

func HomeModel() (Response, error) {
	var data []Home
	var res Response
	cond := db.NewData()
	query := "SHOW TABLES"
	rows, err := cond.Query(query)
	if err != nil {
		log.Println("Query Error")
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var home Home
		err = rows.Scan(&home.TableName)
		if err != nil {
			log.Println("Can't Scan")
			return res, err
		}
		// Tambahkan URL statis ke setiap entri
		home.URL = fmt.Sprintf("http://localhost:1324/%s", home.TableName)
		data = append(data, home)
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res, nil
}
