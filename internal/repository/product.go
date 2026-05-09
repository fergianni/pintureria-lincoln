package repository

import (
	"database/sql"
	"pintureria-lincoln/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	rows, err := r.db.Query(`
		SELECT p.id, p.category_id, c.name, p.name, COALESCE(p.description,''),
		       p.price, COALESCE(p.image_url,''), p.stock, p.active
		FROM products p
		LEFT JOIN categories c ON c.id = p.category_id
		WHERE p.active = TRUE
		ORDER BY p.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.CategoryID, &p.Category, &p.Name,
			&p.Description, &p.Price, &p.ImageURL, &p.Stock, &p.Active); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	var p model.Product
	err := r.db.QueryRow(`
		SELECT p.id, p.category_id, c.name, p.name, COALESCE(p.description,''),
		       p.price, COALESCE(p.image_url,''), p.stock, p.active
		FROM products p
		LEFT JOIN categories c ON c.id = p.category_id
		WHERE p.id = $1 AND p.active = TRUE
	`, id).Scan(&p.ID, &p.CategoryID, &p.Category, &p.Name,
		&p.Description, &p.Price, &p.ImageURL, &p.Stock, &p.Active)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) GetCategories() ([]model.Category, error) {
	rows, err := r.db.Query(`SELECT id, name, slug FROM categories ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug); err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, nil
}
