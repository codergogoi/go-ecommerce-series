package domain

import "time"

type Payment struct {
	ID            uint          `gorm:"PrimaryKey" json:"id"`
	UserId        uint          `json:"user_id"`
	CaptureMethod string        `json:"capture_method"`
	Amount        float64       `json:"amount"`
	OrderId       string        `json:"order_id"`
	CustomerId    string        `json:"customer_id"` // stripe customer if
	PaymentId     string        `json:"payment_id"`  // payment id
	ClientSecret  string        `json:"client_secret"`
	Status        PaymentStatus `json:"status" gorm:"default:initial"` // initial, success, failed
	Response      string        `json:"response"`
	CreatedAt     time.Time     `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time     `gorm:"default:current_timestamp"`
}

type PaymentStatus string

const (
	PaymentStatusInitial PaymentStatus = "initial"
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusPending PaymentStatus = "pending"
)
