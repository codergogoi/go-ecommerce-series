package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignup struct {
	UserLogin
	Phone string `json:"phone"`
}

type VerificationCodeInput struct {
	Code string `json:"code"`
}

type SellerInput struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	BankAccountNumber uint   `json:"bankAccountNumber"`
	SwiftCode         string `json:"swiftCode"`
	PaymentType       string `json:"paymentType"`
}

type AddressInput struct {
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	PostCode     uint   `json:"post_code"`
	Country      string `json:"country"`
}

type ProfileInput struct {
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	AddressInput AddressInput `json:"address"`
}
