package service

type Shipping_Addresses struct {
	Id             int    `json:"id"`
	User_Id        int    `json:"user_id"`
	Recipient_Name string `json:"recipient_name"`
	Addresses      string `json:"addresses"`
	City           string `json:"city"`
	Postal_Code    string `json:"postal_code"`
	Country        string `json:"country"`
	Phone_Number   string `json:"phone_number"`
}

// func Shippping_Addresses()(Response,error) {

// }
