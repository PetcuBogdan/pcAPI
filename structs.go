package main

import "time"


type jsonResponse struct {
	Error 	bool 	`json:"error"`
	Message string 	`json:"message"`
}

type ShopJson struct {
	Shop	string	`json:"shop"`
}

type ProductSupremeJson struct {
	Name 		string	`json:"productName"`
	Category 	string	`json:"category"`
	Colors		string  `json:"colors"`
} 

type ProductPalaceJson struct{
	Name 		string	`json:"productName"`
	Category 	string	`json:"category"`
}

type IdCard struct{
	Id	int	 	`json:"id"`
}

type IdProfile struct{
	Id	int	 	`json:"id"`
}

type IdTask struct{
	Id	int	 	`json:"id"`
}

type Card struct {
	Card_id			int64
    CardNumber 		string
	CardName		string
	CardOwner		string
	Number			int 
    ExpirationDate 	string
    Cvv 			string
	User_id			int64
}

type CardJson struct {
	Card_id			int64    	`json:"id"`
	CardNumber 		string   	`json:"cardNumber"`
    CardName 		string 		`json:"cardName"`
	CardOwner		string 		`json:"cardOwner"`
	Number			int    		`json:"number"`
    ExpirationDate 	string	 	`json:"expirationDate"`
    Cvv 			string		`json:"cvv"`
	UserId			int			`json:"userId"`
}

type Shipment struct {
	Shipment_id 	int64
	User_id			int64
	ShipmentName	string
	Name      		string
	Surname	   		string	 
	Email      		string
	Phone      		string
	Address    		string
	Address2   		string
	Zip   			string
	City       		string
	County	   		string
	Country    		string
}

type ShipmentJson struct {
	Shipment_id		int64		`json:"id"`
	User_id			int64 		`json:"userId"`
	ShipmentName	string		`json:"shipmentName"`
	Name      		string		`json:"name"`
	Surname	   		string		`json:"Surname"`
	Email      		string		`json:"email"`
	Phone      		string		`json:"phone"`
	Address    		string		`json:"address"`
	Address2   		string		`json:"address2"`
	Zip   			string		`json:"zip"`
	City       		string		`json:"city"`
	County	   		string		`json:"county"`
	Country    		string		`json:"country"`
}


type Task struct {
	Shipment_id 	int64
	Card_id 		int64
	Task_id 		int64
	User_id			int64
	Task_name 		string
	Shop 			string
	ProductName 	string
	Category 		string
	Size 			string
	Color 			string
	Status 			string
}



type TaskJson struct {
	Task_id 		int64			`json:"id"`
	User_id			int64			`json:"userId"`
	Task_name 			string			`json:"name"`
	Shop			string			`json:"shop"`
	Card_id 		int64			`json:"card"`
	Shipment_id 	int64			`json:"profile"`
	ProductName 	string			`json:"productName"`
	Category 		string			`json:"category"`
	Size 			string			`json:"size"`
	Color 			string			`json:"color"`
	Status 			string			`json:"status"`
}

type ReadyTask struct {
	Task 		Task
	Shipment	Shipment
	Card 		Card
}

type UserJson struct {
	Id				int64		`json:"userId"`
	Username 		string		`json:"username"`
	Avatar			string		`json:"avatar"`	
	License_id 		int			`json:"licenseId"`
}

type LicenseJson struct {
	License_id 		int64		`json:"licenseId"`
	License_key 	string		`json:"licenseKey"`
	Type 			string		`json:"type"`
	Expiration_date  time.Time	`json:"expirationDate"`
	Creation_date 	time.Time	`json:"creationDate"`
}

type CardID struct {
	Id	int `json:"id"`
}

type ShipmentID struct{
	Id int `json:"id"`
}

type TaskID struct {
	ID int `json:"id"`
}

type UserID struct {
	ID int `json:"id"`
}

type LicenseKey struct {
	License_key 	string		`json:"license"`
}

type LicenseVerification struct {
	License_key 	string		`json:"license"`
	User_id			int			`json:"userId"`
}

