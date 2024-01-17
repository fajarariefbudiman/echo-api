package service

import (
	"database/sql"
	"echo-api/db"
	"log"
	"net/http"
)

type Orders struct {
	Id                 int `json:"id"`
	User_Id            int `json:"user_id"`
	Shipping_Addresses Shipping_Addresses
	Order_Status       string        `json:"order_status"`
	Total_Price        float64       `json:"total_price"`
	Payment_Method     string        `json:"payment_method"`
	Order_Date         string        `json:"order_date"`
	Order_Items        []Order_Items `json:"order_items"`
}

func AddOrders(User_Id int, Order_Status string, Payment_Method string) (Response, error) {
	var res Response
	cond := db.NewData()
	query := "INSERT INTO orders (User_Id,Order_Status,Payment_Method) VALUES (?,?,?)"
	sqlstmnt, err := cond.Prepare(query)
	if err != nil {
		log.Println("Query Prepare Error")
		return res, err
	}
	result, err := sqlstmnt.Exec(&User_Id, &Order_Status, &Payment_Method)
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

func UpdateOrders(Order_Id int) (Response, error) {
	var res Response
	cond := db.NewData()
	query := `
		UPDATE orders
		SET Total_Price = (
			SELECT SUM(Product_Price * Total_Product) 
			FROM order_items 
			WHERE Order_Id = ?
		)
		WHERE Id = ?
	`

	sqlstmnt, err := cond.Prepare(query)
	if err != nil {
		log.Println("Query Prepare Error")
		return res, err
	}

	result, err := sqlstmnt.Exec(&Order_Id, &Order_Id)
	if err != nil {
		log.Println("Exec Error")
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("No result because %s", err)
		return res, err
	}

	if rowsAffected == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Order not found"
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Success Update"
	return res, nil
}

func GetOrderById(Id int) (Response, error) {
	var res Response
	var order Orders
	cond := db.NewData()
	query := "SELECT * FROM orders WHERE id=?"
	row, err := cond.Query(query, Id)
	if err != nil {
		log.Println("Query Error No Row")
		return res, err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&order.Id, &order.User_Id, &order.Order_Status, &order.Total_Price, &order.Payment_Method, &order.Order_Date)
		if err == sql.ErrNoRows {
			return res, sql.ErrNoRows
		}
	}
	var order_items []Order_Items
	queryitems := "SELECT * FROM order_items WHERE order_id=?"
	rows, err := cond.Query(queryitems, order.Id)
	if err != nil {
		log.Println("Row Cart_Items Error")
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var order_item Order_Items
		err := rows.Scan(&order_item.Id, &order_item.Product_Id, &order_item.Order_Id,
			&order_item.Product_Name, &order_item.Product_Price, &order_item.Total_Product)
		if err != nil {
			log.Println("Rows Scan Error")
			return res, err
		}
		order_items = append(order_items, order_item)
	}
	order.Order_Items = order_items
	res.Status = http.StatusOK
	res.Message = "Success Get order"
	res.Data = order

	return res, nil
}
