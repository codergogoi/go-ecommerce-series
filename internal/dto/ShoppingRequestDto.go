package dto

type CreateCartRequest struct {
	ProductId uint `json:"product_id"`
	Qty       uint `json:"qty"`
}
