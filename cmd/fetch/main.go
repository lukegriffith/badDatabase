package main

import (
	"fmt"
	"log"
	"github.com/lukegriffith/badDatabase/internal/ecommerce/database"
	"github.com/lukegriffith/badDatabase/internal/ecommerce"
)

func main() {
	dsn := "postgres://baddb:test@localhost/ecommerce"
	db, err := database.Connect(dsn)
	if err != nil {
		log.Panicln("Error connecting to database:", err)
		return
	}
	orders, err := ecommerce.FetchAllData(db)
	if err != nil {
		log.Panicln(err)
	}
	for _, order := range orders {
		fmt.Printf("Order ID: %s\n", order.OrderID)
		fmt.Printf("Customer Name: %s\n", order.Customer.Name)
		fmt.Printf("Customer Email: %s\n", order.Customer.Email)

		for _, orderItem := range order.OrderItems {
			fmt.Printf("\tOrder Item ID: %s\n", orderItem.OrderItemID)
			fmt.Printf("\tQuantity: %d\n", orderItem.Quantity)
			fmt.Printf("\tProduct Name: %s\n", orderItem.Product.ProductName)
			fmt.Println()
		}
		fmt.Println()
	}

}

