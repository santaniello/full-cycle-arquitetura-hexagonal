package main

import (
	"database/sql"
	db2 "github.com/santaniello/full-cycle-arquitetura-hexagonal/adapters/db"
	"github.com/santaniello/full-cycle-arquitetura-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productAdapter)
	product, _ := productService.Create("Product Exemplo", 30)
	productService.Enable(product)
}
