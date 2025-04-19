package dto

type CreateCartRequest struct {
	ProductId uint `json:"product_id"`
	Qty       uint `json:"qty"`
}

type CreatePaymentRequest struct {
	OrderId      string  `json:"order_id"`
	PaymentId    string  `json:"payment_id"`
	ClientSecret string  `json:"client"`
	Amount       float64 `json:"amount"`
	UserId       uint    `json:"user_id"`
}
