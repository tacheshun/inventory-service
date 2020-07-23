package main

import (
	"inventoryService/database"
	"inventoryService/product"
	"inventoryService/receipt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}