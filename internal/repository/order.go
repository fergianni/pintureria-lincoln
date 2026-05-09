package repository

import (
	"database/sql"
	"pintureria-lincoln/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(req model.CreateOrderRequest, total float64) (*model.Order, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var order model.Order
	err = tx.QueryRow(`
		INSERT INTO orders (name, email, phone, notes, total)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, email, phone, COALESCE(notes,''), total, status, created_at
	`, req.Name, req.Email, req.Phone, req.Notes, total).
		Scan(&order.ID, &order.Name, &order.Email, &order.Phone,
			&order.Notes, &order.Total, &order.Status, &order.CreatedAt)
	if err != nil {
		return nil, err
	}

	for _, item := range req.Items {
		_, err = tx.Exec(`
			INSERT INTO order_items (order_id, product_id, quantity, unit_price)
			VALUES ($1, $2, $3, $4)
		`, order.ID, item.ProductID, item.Quantity, item.UnitPrice)
		if err != nil {
			return nil, err
		}
	}

	order.Items = req.Items
	return &order, tx.Commit()
}
