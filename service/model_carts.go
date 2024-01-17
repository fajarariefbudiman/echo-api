package service

import (
	"database/sql"
	"echo-api/db"
	"log"
	"net/http"
)

type Carts struct {
	Id         int
	User_Id    int
	Cart_Items []Cart_Items
}

func AddCarts(User_Id int) (Response, error) {
	var res Response
	cond := db.NewData()
	query := "INSERT INTO user_cart (User_Id) VALUES (?)"
	sqlstmt, err := cond.Prepare(query)
	if err != nil {
		log.Println("Query Error")
		return res, err
	}
	result, err := sqlstmt.Exec(&User_Id)
	if err != nil {
		log.Println("No result")
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "Success Add Cart_Items"
	res.Data = result

	return res, nil
}

func GetCartByUserId(Id int) (Response, error) {
	var res Response
	var cart Carts
	cond := db.NewData()
	query := "SELECT Id,User_Id FROM user_cart WHERE user_id=?"
	row, err := cond.Query(query, Id)
	if err != nil {
		log.Println("Query Error No Row")
		return res, err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&cart.Id, &cart.User_Id)
		if err == sql.ErrNoRows {
			return res, sql.ErrNoRows
		}
	}
	var cart_items []Cart_Items
	queryitems := "SELECT cart_id,product_id,quantity,price FROM cart_items WHERE cart_id=?"
	rows, err := cond.Query(queryitems, cart.Id)
	if err != nil {
		log.Println("Row Cart_Items Error")
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var cart_item Cart_Items
		err := rows.Scan(&cart_item.Cart_Id, &cart_item.Product_Id, &cart_item.Quantity,
			&cart_item.Price)
		if err != nil {
			log.Println("Rows Scan Error")
			return res, err
		}
		cart_items = append(cart_items, cart_item)
	}

	cart.Cart_Items = cart_items
	res.Status = http.StatusOK
	res.Message = "Success Get order"
	res.Data = cart

	return res, nil
}
