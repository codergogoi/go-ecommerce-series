package repository

import (
	"gorm.io/gorm"
)

type CatalogRepository interface {
}

type catalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}
