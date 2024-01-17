package service

import (
	"echo-api/db"
	"log"
	"net/http"
)

type Cart_Items struct {
	Cart_Id    int
	Product_Id int
	Quantity   int
	Price      float32
	Products   []Products
}

func AddCartItems(Cart_Id, Product_Id, Quantity int, Price float32) (Cart_Items, error) {
	var cart Cart_Items
	cond := db.NewData()
	query := "INSERT INTO cart_items (Cart_Id, Product_Id, Quantity, Price) VALUES (?, ?, ?, ?)"
	sqlstmt, err := cond.Prepare(query)
	if err != nil {
		log.Println("Query Error")
		return cart, err
	}
	_, err = sqlstmt.Exec(&Cart_Id, &Product_Id, &Quantity, &Price)
	if err != nil {
		log.Println("Error executing SQL statement")
		return cart, err
	}

	// // Retrieve product details
	// productquery := "SELECT id, name, price, stock, discount FROM products WHERE id=?"
	// productRow := cond.QueryRow(productquery, Product_Id)

	// var product Products
	// if err := productRow.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &product.Discount); err != nil {
	// 	log.Println("Error fetching product details:", err)
	// 	return cart, err
	// }

	// // Populate the cart variable
	cart.Cart_Id = Cart_Id
	cart.Product_Id = Product_Id
	cart.Quantity = Quantity
	cart.Price = Price
	// cart.Product = product

	return cart, nil
}

func GetCartItems(Cart_Id int) (Response, error) {
	var res Response
	var cart_items Cart_Items
	cond := db.NewData()
	query := "SELECT cart_id,product_id,quantity,price FROM cart_items WHERE cart_id = ?"
	row, err := cond.Query(query, Cart_Id)
	if err != nil {
		log.Println("Query Error")
		return res, err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&cart_items.Cart_Id, &cart_items.Product_Id, &cart_items.Quantity, &cart_items.Price)
		if err != nil {
			log.Println("No Rows Selected")
			return res, err
		}
	}

	var products []Products
	queryproducts := "SELECT id,name,slug,category_id,price,discount FROM products WHERE id = ?"
	rows, err := cond.Query(queryproducts, cart_items.Product_Id)
	if err != nil {
		log.Println("Query Products Error")
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Products
		err = rows.Scan(&product.Id, &product.Name, &product.Slug, &product.Category_Id, &product.Price, &product.Discount)
		if err != nil {
			log.Println("No Rows Selected In Table Products")
			return res, err
		}
		products = append(products, product)
	}
	cart_items.Products = products
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = cart_items

	return res, nil
}
