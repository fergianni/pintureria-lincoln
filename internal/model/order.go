package model

import "time"

type OrderItem struct {
	ProductID int     `json:"product_id"`
	Name      string  `json:"name,omitempty"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type Order struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Notes     string      `json:"notes"`
	Items     []OrderItem `json:"items"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

type CreateOrderRequest struct {
	Name  string      `json:"name"  binding:"required"`
	Email string      `json:"email" binding:"required,email"`
	Phone string      `json:"phone"`
	Notes string      `json:"notes"`
	Items []OrderItem `json:"items" binding:"required,min=1"`
}
