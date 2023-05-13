package database 

import (
	"time"
	"github.com/google/uuid"
)

type Product struct {
	ProductID   uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk"`
	ProductName string    `bun:",notnull"`
	Description string
	Price       float64 `bun:",notnull"`
	Category    string  `bun:",notnull"`
}

type Customer struct {
	CustomerID  uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk"`
	Name        string    `bun:",notnull"`
	Email       string    `bun:",unique,notnull"`
	PhoneNumber string
}

type Order struct {
	OrderID    uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk"`
	CustomerID uuid.UUID  `bun:"type:uuid,notnull"`
	OrderDate  time.Time `bun:",notnull"`

	Customer *Customer `bun:"rel:belongs-to,join:customer_id=customer_id"` // Add this line
	OrderItems []*OrderItem `bun:"rel:has-many,join:order_id=order_id"`
}

type OrderItem struct {
	OrderItemID uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk"`
	OrderID     uuid.UUID `bun:"type:uuid,notnull"`
	ProductID   uuid.UUID `bun:"type:uuid,notnull"`
	Quantity    int       `bun:",notnull"`
	TotalPrice  float64   `bun:",notnull"`

	Order   *Order   `bun:"rel:belongs-to,join:order_id=order_id"` // Add this line
	Product *Product `bun:"rel:belongs-to,join:product_id=product_id"`
}
