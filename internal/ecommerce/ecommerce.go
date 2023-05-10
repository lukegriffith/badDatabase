package ecommerce

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/lukegriffith/badDatabase/internal/ecommerce/database"
	"github.com/bxcodec/faker/v3"
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

	// Insert fake products
	for i := 0; i < numProducts; i++ {
		product := database.Product{
			Name:        faker.Name(),
			Description: faker.Sentence(),
			Price:       faker.Float64Range(1, 100),
		}
		_, err := db.NewInsert().Model(&product).Exec(ctx)
		if err != nil {
			return fmt.Errorf("error inserting product: %w", err)
		}
	}

	// Insert fake customers
	for i := 0; i < numCustomers; i++ {
		customer := database.Customer{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
		}
		_, err := db.NewInsert().Model(&customer).Exec(ctx)
		if err != nil {
			return fmt.Errorf("error inserting customer: %w", err)
		}
	}

	// Insert fake orders and order items
	for i := 0; i < numOrderItems; i++ {
		orders := database.OrderItem{
			// TODO	
		}
	}



