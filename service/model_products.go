package service

import (
	"database/sql"
	"echo-api/db"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

type Products struct {
	Id          int64   `json:"id"`
	Name        string  `validate:"required" json:"name"`
	Slug        string  `validate:"required" json:"slug"`
	Category_Id int     `validate:"required" json:"category_id"`
	Price       float32 `validate:"required" json:"price"`
	Discount    float32 `json:"discount"`
	Stock       int     `validate:"required" json:"stock"`
	Created_At  string  `json:"created_at"`
	Updated_At  string  `json:"updated_at"`
}

func GetAllProducts() (Response, error) {
	var get Products
	var getAll []Products
	var res Response

	cond := db.NewData()
	sqlstmt := "SELECT products.* FROM products INNER JOIN categories ON products.category_id = categories.id"
	rows, err := cond.Query(sqlstmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Name, &get.Slug, &get.Category_Id, &get.Price, &get.Discount, &get.Stock, &get.Created_At, &get.Updated_At)
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

func CreateProducts(name string, slug string, price, discount float32, stock int, category_id int) (Response, error) {
	var response Response
	v := validator.New()
	Use := Products{
		Name:        name,
		Slug:        slug,
		Price:       price,
		Stock:       stock,
		Category_Id: category_id,
	}
	err := v.Struct(Use)
	if err != nil {
		return response, err
	}
	cond := db.NewData()
	sqlstmnt := "INSERT INTO products (name,slug,price,discount,stock,category_id) VALUES(?,?,?,?,?,?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(name, slug, price, discount, stock, category_id)
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

func DeleteProducts(Slug string) (Response, error) {
	var response Response
	cond := db.NewData()
	sqlstmnt := "DELETE FROM products WHERE slug=?"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}

	result, err := stmt.Exec(Slug)
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

func UpdateProducts(Slug string, Name string, Price int, Stock int, Category_Id int) (Response, error) {
	var response Response
	cond := db.NewData()
	sqlstmnt := "UPDATE products SET name=?, price=?, stock=?, category_id=? WHERE slug=?"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(Name, Price, Stock, Category_Id, Slug)
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

func GetProductsSlug(Slug string) (Response, error) {
	var get Products
	var res Response
	cond := db.NewData()
	sqlstmt := "SELECT id,slug,name,category_id,price,discount,stock,created_at,updated_at FROM products WHERE slug=?"
	rows, err := cond.Query(sqlstmt, Slug)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Slug, &get.Name, &get.Category_Id, &get.Price, &get.Discount, &get.Stock, &get.Created_At, &get.Updated_At)
		if err == sql.ErrNoRows {
			return res, sql.ErrNoRows
		}
		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = get

	}

	return res, nil
}
