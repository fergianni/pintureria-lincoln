package model

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Product struct {
	ID          int      `json:"id"`
	CategoryID  int      `json:"category_id"`
	Category    string   `json:"category,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	ImageURL    string   `json:"image_url"`
	Stock       int      `json:"stock"`
	Active      bool     `json:"active"`
}
