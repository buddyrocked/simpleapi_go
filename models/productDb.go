package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetProducts() []Product {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM product")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	products := []Product{}
	for results.Next() {
		var prod Product

		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)

		if err != nil {
			panic(err.Error())
		}

		products = append(products, prod)
	}

	return products
}

func GetProduct(code string) *Product {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	prod := &Product{}

	if err != nil {
		fmt.Println("Err", err.Error())

		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM product WHERE code = ?", code)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return prod
}

func AddProduct(product Product) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO product (code, name, qty, last_updated) VALUES (?, ?, ?, now())", product.Code, product.Name, product.Qty)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func UpdateProduct(code string, product Product) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	update, err := db.Query("UPDATE product set code = ?, name = ?,  qty = ?, last_updated = now() WHERE code = ?", product.Code, product.Name, product.Qty, product.Code)

	if err != nil {
		panic(err.Error())
	}

	defer update.Close()
}

func DeleteProduct(code string) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	delete, err := db.Query("DELETE FROM product WHERE code = ?", code)

	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()
}
