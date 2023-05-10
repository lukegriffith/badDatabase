package main

import (
	"fmt"
	"github.com/lukegriffith/badDatabase/internal/ecommerce/database"
	"github.com/lukegriffith/badDatabase/internal/ecommerce"
)

func main() {
	dsn := "postgres://baddb:test@localhost/ecommerce"
	db, err := database.Connect(dsn)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// Perform operations using the db instance
	err = ecommerce.CreateTables(db)
	if err != nil {
		panic(err)
	}
}

