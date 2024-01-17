package service

import (
	"database/sql"
	"echo-api/db"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

type Categories struct {
	Id         int64  `json:"id"`
	Name       string `validate:"required" json:"name"`
	Slug       string `validate:"required" json:"slug"`
	Created_At string `json:"created_at"`
	Products   []Products
}

func GetAllCategories() (Response, error) {
	var categories []Categories
	var res Response
	cond := db.NewData()
	sqlstmt := "SELECT * FROM categories"
	rows, err := cond.Query(sqlstmt)
	if err != nil {
		log.Println("No Query Rows")
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category Categories
		err := rows.Scan(&category.Id, &category.Name, &category.Slug, &category.Created_At)
		if err != nil {
			log.Printf("No Rows: %s", err)
		}
		categories = append(categories, category)
	}

	for i := range categories {
		var products []Products
		query := "SELECT * FROM products WHERE category_id = ?"
		rows, err := cond.Query(query, categories[i].Id)
		if err != nil {
			log.Printf("No rows: %s", err)
		}
		defer rows.Close()

		for rows.Next() {
			var product Products
			err := rows.Scan(&product.Id, &product.Name, &product.Slug, &product.Category_Id, &product.Price, &product.Discount,
				&product.Stock, &product.Created_At, &product.Updated_At)
			if err != nil {
				log.Printf("Error rows Products: %s", err)
			}
			products = append(products, product)
		}

		categories[i].Products = products

		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = categories

	}
	return res, nil
}
func CreateCategories(name string, slug string) (Response, error) {
	var response Response
	v := validator.New()
	Use := Categories{
		Name: name,
		Slug: slug,
	}
	err := v.Struct(Use)
	if err != nil {
		return response, err
	}
	cond := db.NewData()
	sqlstmnt := "INSERT INTO categories (name,slug) VALUES(?,?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(name, slug)
	if err != nil {
		return response, err
	}
	lastinsert, err := result.LastInsertId()
	if err != nil {
		log.Printf("No result because %s", err)
		return response, err
	}
	response.Status = http.StatusCreated
	response.Message = "Success Create"
	response.Data = map[string]int64{
		"lastinsertId": lastinsert,
	}
	return response, err
}

func DeleteCategories(Id int) (Response, error) {
	var response Response
	cond := db.NewData()
	sqlstmnt := "DELETE FROM categories WHERE id=?"
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

func UpdateCategories(Id int, Name string, Slug string) (Response, error) {
	var response Response
	cond := db.NewData()
	sqlstmnt := "UPDATE categories set name=?, slug=? WHERE id=?"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(Name, Slug, Id)
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

func GetCategoriesSlug(Slug string) (Response, error) {
	var get Categories
	var res Response
	cond := db.NewData()
	sqlstmt := "SELECT id,name,slug, created_at FROM categories WHERE slug=?"
	rows, err := cond.Query(sqlstmt, Slug)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Name, &get.Slug, &get.Created_At)
		if err == sql.ErrNoRows {
			return res, sql.ErrNoRows
		}

	}
	var products []Products
	query := "SELECT * FROM products WHERE category_id = ?"
	rows, err = cond.Query(query, get.Id)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Products
		err := rows.Scan(&product.Id, &product.Name, &product.Slug, &product.Category_Id, &product.Price, &product.Discount,
			&product.Stock, &product.Created_At, &product.Updated_At)
		if err != nil {
			log.Printf("Error rows Products: %s", err)
		}
		products = append(products, product)
	}

	get.Products = products
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = get
	return res, nil
}
