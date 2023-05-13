package ecommerce

import (
	"context"
	"fmt"
	"reflect"
	"time"
	"math/rand"

	faker "github.com/bxcodec/faker/v3"
	"github.com/lukegriffith/badDatabase/internal/ecommerce/database"
	"github.com/uptrace/bun"
)

func CreateTables(db *bun.DB) error {
	models := []interface{}{
		(*database.Product)(nil),
		(*database.Customer)(nil),
		(*database.Order)(nil),
		(*database.OrderItem)(nil),
	}

	for _, model := range models {
		_, err := db.NewDropTable().Model(model).IfExists().Exec(context.Background())
		if err != nil {
			return fmt.Errorf("error dropping table: %w", err)
		}

		_, err = db.NewCreateTable().Model(model).Exec(context.Background())
		if err != nil {
			return fmt.Errorf("error creating table: %w", err)
		}
	}

	return nil
}

// Add more functions for CRUD operations and other database-related tasks

func InsertFakeData(db *bun.DB, numProducts, numCustomers, numOrders, numOrderItems int) error {
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())
	// Insert fake products
	for i := 0; i < numProducts; i++ {
		var price float64
		v := reflect.ValueOf(&price).Elem()
		p, err := faker.GetPrice().Amount(v)
		price = p.(float64)
		if err != nil {
			panic(err)
		}
		product := database.Product{
			ProductName:        faker.Name(),
			Description: faker.Sentence(),
			Price: price,
		}
		_, err = db.NewInsert().Model(&product).Exec(ctx)
		if err != nil {
			return fmt.Errorf("error inserting product: %w", err)
		}
	}

	// Insert fake customers
	for i := 0; i < numCustomers; i++ {
		customer := database.Customer{
			Name: fmt.Sprintf("%s %s", faker.FirstName(), faker.LastName()),
			Email:     faker.Email(),
			PhoneNumber: faker.Phonenumber(),
		}
		_, err := db.NewInsert().Model(&customer).Exec(ctx)
		if err != nil {
			return fmt.Errorf("error inserting customer: %w", err)
		}
	}

	// Retrieve all customers
	var customers []database.Customer
	err := db.NewSelect().Model(&customers).Scan(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving customers: %w", err)
	}

	// Check if there are customers available
	if len(customers) == 0 {
		return fmt.Errorf("error: no customers found")
	}



	// Insert fake orders
	for i := 0; i < numOrders; i++ {
		order := database.Order{
			CustomerID: customers[rand.Intn(len(customers))].CustomerID,
			OrderDate: time.Now(),
		}
		_, err := db.NewInsert().Model(&order).Exec(ctx)
		if err != nil {
			return fmt.Errorf("error inserting order: %w", err)
		}
	}


	// Retrieve all orders and products
	var orders []database.Order
	err = db.NewSelect().Model(&orders).Scan(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving orders: %w", err)
	}

	var products []database.Product
	err = db.NewSelect().Model(&products).Scan(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving products: %w", err)
	}

	// Check if there are orders and products available
	if len(orders) == 0 || len(products) == 0 {
		return fmt.Errorf("error: no orders or products found")
	}



	// Insert fake order items
	for i := 0; i < numOrderItems; i++ {
		var price float64
		v := reflect.ValueOf(&price).Elem()
		p, err := faker.GetPrice().Amount(v)
		price = p.(float64)
		orderItem := database.OrderItem{
			OrderID:   orders[rand.Intn(len(orders))].OrderID,
			ProductID: products[rand.Intn(len(products))].ProductID,
			Quantity:  rand.Intn(10) + 1,
			TotalPrice:     price,
		}
		_, err = db.NewInsert().Model(&orderItem).Exec(ctx)
		if err != nil {
			return fmt.Errorf("error inserting order item: %w", err)
		}
	}
	return nil
}

