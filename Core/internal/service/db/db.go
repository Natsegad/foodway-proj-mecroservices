package db

import (
	"core/internal/domain"
	"core/pkg/db"
)

func AutoMigrate() {
	db.DataBase.AutoMigrate(&domain.Product{})
}
