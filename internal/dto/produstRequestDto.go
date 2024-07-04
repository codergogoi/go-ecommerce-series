package dto

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryId  uint    `json:"category_id"`
	ImageUrl    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type UpdateStockRequest struct {
	Stock int `json:"stock"`
}
