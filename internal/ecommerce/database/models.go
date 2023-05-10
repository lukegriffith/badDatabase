package database 

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ProductID   uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	ProductName string    `bun:",notnull"`
	Description string
	Price       float64 `bun:",notnull"`
	Category    string  `bun:",notnull"`
}

type Customer struct {
	CustomerID  uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Name        string    `bun:",notnull"`
	Email       string    `bun:",unique,notnull"`
	PhoneNumber string
}

type Order struct {
	OrderID    uuid.UUID  `bun:"type:uuid,default:uuid_generate_v4()"`
	CustomerID uuid.UUID  `bun:",notnull"`
	OrderDate  time.Time `bun:",notnull"`
}

type OrderItem struct {
	OrderItemID uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	OrderID     uuid.UUID `bun:",notnull"`
	ProductID   uuid.UUID `bun:",notnull"`
	Quantity    int       `bun:",notnull"`
	TotalPrice  float64   `bun:",notnull"`
}
