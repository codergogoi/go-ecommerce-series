package dto

type SellerOrderDetails struct {
	OrderRefNumber  int    `json:"order_ref_number"`
	OrderStatus     int    `json:"order_status"`
	CreatedAt       string `json:"created_at"`
	OrderItemId     uint   `json:"order_item_id"`
	ProductId       uint   `json:"product_id"`
	Name            string `json:"name"`
	ImageUrl        string `json:"image_url"`
	Price           string `json:"price"`
	Qty             uint   `json:"qty"`
	CustomerName    string `json:"customer_name"`
	CustomerEmail   string `json:"customer_email"`
	CustomerPhone   string `json:"customer_phone"`
	CustomerAddress string `json:"customer_address"`
}
