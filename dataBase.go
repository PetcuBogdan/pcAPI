package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


const(
	host="localhost"
	port=5432
	user="postgres"
	password="Petcu123"
	dbname="ckt"
)


func saveCard(card CardJson) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}

	insertStmt := `insert into creditcards(card_number, card_name, card_owner, number, expiration_date, cvv) values(?, ?, ?, ?, ?, ?)`
    db.Exec(insertStmt, card.CardNumber, card.CardName, card.CardOwner, card.Number, card.ExpirationDate, card.Cvv)
	
	defer func() {
		dbInstance, _ := db.DB()
    _ = dbInstance.Close()
	}()
}

func editCard(card CardJson) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}

	editStmt := `update creditcards set card_number = ?, card_name = ?, card_owner = ?, number = ?, expiration_date = ?, cvv = ? where card_id = ?`
    db.Exec(editStmt, card.CardNumber, card.CardName, card.CardOwner, card.Number, card.ExpirationDate, card.Cvv, card.Card_id)

}

func deleteCard(cardId int) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}

	deleteStmt := `delete from "creditcards" where card_id=?`
    db.Exec(deleteStmt, cardId)
}

func getListOfCards() []CardJson {
		dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
		db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
		if err != nil{
			log.Fatal(err)
		}else{
			fmt.Printf("Conexiune reusita")
		}
		
		var cards []CardJson
		db.Raw(`select * from creditcards`).Scan(&cards)
	
	return cards;
}

func saveShipment(shipment ShipmentJson){
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
		db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
		if err != nil{
			log.Fatal(err)
		}else{
			fmt.Printf("Conexiune reusita")
		}

		saveStmt := `insert into shipment(shipment_name, name, surname, address, address_extend, city, county, country, zip, email, phone) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		db.Exec(saveStmt, shipment.ShipmentName, shipment.Name, shipment.Surname, shipment.Address, shipment.Address2, shipment.City, shipment.County, shipment.Country, shipment.Zip, shipment.Email, shipment.Phone)
		
}

func editShipment(shipment ShipmentJson) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
		db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
		if err != nil{
			log.Fatal(err)
		}else{
			fmt.Printf("Conexiune reusita")
		}

		editStmt := `update shipment set shipment_name=?, name=?, surname=?, address=?, address_extend=?, city=?, county=?, country=?, zip=?, email=?, phone=? where shipment_id = ?`
		db.Exec(editStmt, shipment.ShipmentName, shipment.Name, shipment.Surname, shipment.Address, shipment.Address2, shipment.City, shipment.County, shipment.Country, shipment.Zip, shipment.Email, shipment.Phone, shipment.Shipment_id)
		
}

func deleteShipment(shipmentId int) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}

	deleteStmt := `delete from shipment where shipment_id=?`
    db.Exec(deleteStmt, shipmentId)
}

func getListOfShipments() []ShipmentJson {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
		db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
		if err != nil{
			log.Fatal(err)
		}else{
			fmt.Printf("Conexiune reusita")
		}
		
		var shipments []ShipmentJson
		db.Raw(`select * from shipment`).Scan(&shipments)
	
	return shipments;
}

func saveTask(task TaskJson) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
		db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
		if err != nil{
			log.Fatal(err)
		}else{
			fmt.Printf("Conexiune reusita")
		}

		insertStmt := `insert into tasks(task_name, website, product_name, size, color, category, creditcard_id, shipment_id, status, user_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		db.Exec(insertStmt, task.Task_name, task.Shop, task.ProductName, task.Size, task.Color, task.Category, task.Card_id, task.Shipment_id, task.Status, task.User_id)
		
}

func editTask(task TaskJson) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
		db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
		if err != nil{
			log.Fatal(err)
		}else{
			fmt.Printf("Conexiune reusita")
		}

		editStmt := `update tasks set task_name=?, website=?, product_name=?, size=?, color=?, category=?, creditcard_id=?, shipment_id=?, status=? where task_id = ?`
		db.Exec(editStmt, task.Task_name, task.Shop, task.ProductName, task.Size, task.Color, task.Category, task.Card_id, task.Shipment_id, task.Status, task.Task_id)		
}

func deleteTask(taskId int) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}

	deleteStmt := `delete from tasks where task_id=?`
    db.Exec(deleteStmt, taskId)
}

func getListOfTasks() []TaskJson {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
		db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
		if err != nil{
			log.Fatal(err)
		}else{
			fmt.Printf("Conexiune reusita")
		}
		
		var tasks []TaskJson
		db.Raw(`select * from tasks`).Scan(&tasks)
	return tasks;
}

func getSpecTask(id int) Task{
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}
		
	var tasks []Task
	db.Raw(`select * from tasks`).Scan(&tasks)
	var taskUsed Task
	for  _,task := range tasks{
		if int(task.Task_id) == id{
			taskUsed = task
			break
		}
	}
	return taskUsed
}

func getSpecCard(id int) Card{
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}
		
	var cards []Card
	db.Raw(`select * from creditcards`).Scan(&cards)
	var cardUsed Card
	for  _,card := range cards{
		if int(card.Card_id) == id{
			cardUsed = card
			break
		}
	}
	return cardUsed
}

func getSpecShipment(id int) Shipment{
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	
	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}
		
	var shipments []Shipment
	db.Raw(`select * from shipment`).Scan(&shipments)
	var shipmentUsed Shipment
	for  _,shipment := range shipments{
		if int(shipment.Shipment_id) == id{
			shipmentUsed = shipment
			break
		}
	}
	return shipmentUsed
}

func saveUser(appUser UserJson) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}
	
	var users []UserJson
	var ver int
	ver = 0
	db.Raw(`select * from users`).Scan(&users)
	for _,findUser := range users{
		if findUser.Id == appUser.Id {
			ver = 1
		}
	}

	if ver == 0{
		insertStmt := `insert into users(id, username, avatar) values(?, ?, ?)`
    	db.Exec(insertStmt, appUser.Id, appUser.Username, appUser.Avatar)
	}

	defer func() {
		dbInstance, _ := db.DB()
    _ = dbInstance.Close()
	}()
}

func saveLicense(licenseKey string, id int) string {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}
	
	var license LicenseJson

	db.Raw(`select * from licenses where license_key=?`, licenseKey).Scan(&license)
	fmt.Println(license)
	
	var user UserJson
	db.Raw(`select * from users where license_id=?`, license.License_id).Scan(&user)

	fmt.Println(user)

	if user.Id == 0{
		insertLicenseKey := `update users set license_id=? where id=?`
    	db.Exec(insertLicenseKey, license.License_id, id)
		defer func() {
			dbInstance, _ := db.DB()
		_ = dbInstance.Close()
		}()
		return "license saved"
	} else {
		return "license already used"
	} 
}

func deleteLicense(licenseKey string) {
	dbConn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("Conexiune reusita")
	}

	var license LicenseJson
	db.Raw(`select * from licenses where license_key=?`, licenseKey).Scan(&license)
	fmt.Println(license)
	deleteLicense := `update users set license_id = NULL where license_id = ?`
    db.Exec(deleteLicense,license.License_id)

}