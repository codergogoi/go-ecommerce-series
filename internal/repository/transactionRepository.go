package repository

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreatePayment(payment *domain.Payment) error
	FindInitialPayment(uId uint) (*domain.Payment, error)
	UpdatePayment(payment *domain.Payment) error
	FindOrders(uId uint) ([]domain.OrderItem, error)
	FindOrderById(uId uint, id uint) (dto.SellerOrderDetails, error)
}

type transactionStorage struct {
	db *gorm.DB
}

func (t transactionStorage) UpdatePayment(payment *domain.Payment) error {
	return t.db.Save(payment).Error
}

func (t transactionStorage) FindInitialPayment(uId uint) (*domain.Payment, error) {
	var payment *domain.Payment
	err := t.db.First(&payment, "user_id=? AND status=?", uId, "initial").Order("created_at desc").Error
	return payment, err
}

func (t transactionStorage) CreatePayment(payment *domain.Payment) error {
	return t.db.Create(payment).Error
}

func (t transactionStorage) FindOrders(uId uint) ([]domain.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}

func (t transactionStorage) FindOrderById(uId uint, id uint) (dto.SellerOrderDetails, error) {
	//TODO implement me
	panic("implement me")
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionStorage{db: db}
}
