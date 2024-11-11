package repository

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreatePayment(payment *domain.Payment) error
	FindOrders(uId uint) ([]domain.OrderItem, error)
	FindOrderById(uId uint, id uint) (dto.SellerOrderDetails, error)
}

type transactionStorage struct {
	db *gorm.DB
}

func (t transactionStorage) CreatePayment(payment *domain.Payment) error {
	//TODO implement me
	panic("implement me")
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
