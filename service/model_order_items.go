package service

import (
	"echo-api/db"
	"log"
	"net/http"
)

type Order_Items struct {
	Id            int     `json:"id"`
	Order_Id      int     `json:"order_id"`
	Product_Id    int     `json:"product_id"`
	Product_Name  string  `json:"product_name"`
	Product_Price float32 `json:"product_price"`
	Total_Product int     `json:"total_product"`
	Products      Products
}

// func GetOrderItems() {

// }

func AddOrderItems(Order_Id, Product_Id int, Product_Name string, Product_Price float32, Total_Product int) (Response, error) {
	var res Response
	cond := db.NewData()
	query := "INSERT INTO order_items (Order_Id,Product_Id,Product_Name,Product_Price,Total_Product) VALUES (?,?,?,?,?)"
	sqlstmnt, err := cond.Prepare(query)
	if err != nil {
		log.Println("Query Prepare Error")
		return res, err
	}
	result, err := sqlstmnt.Exec(&Order_Id, &Product_Id, &Product_Name, &Product_Price, &Total_Product)
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
