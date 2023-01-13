package models

import (
	"bharvest-vo/types"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetProducts() []Product {
	config := types.GetConfig()
	var (
		DB_USER = config.Mysql.DbUser
		DB_PASS = config.Mysql.DbPass
		DB_NAME = config.Mysql.DbName
	)

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp(127.0.0.1:3306)/"+DB_NAME)

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT address, createdAt FROM Users")
	fmt.Println(results)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	products := []Product{}
	for results.Next() {
		var prod Product
		// for each row, scan into the Product struct
		err = results.Scan(&prod)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the product into products array
		products = append(products, prod)
	}

	return products

}
