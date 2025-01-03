package domain

import "time"

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type User struct {
	ID        uint      `json:"id" gorm:"PrimaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      string    `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Address   Address   `json:"address"` // relation
	Cart      Cart      `json:"cart"`    // relation
	Orders    []Order   `json:"orders"`  // relation
	Payments  []Payment `json:"payment"` // relation
	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"user_type" gorm:"default:buyer"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
