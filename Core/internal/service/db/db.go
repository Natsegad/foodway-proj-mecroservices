package db

import (
	"core/internal/domain"
	"core/pkg/db"
)

func AutoMigrateService() {
	db.DataBase.AutoMigrate(
		&domain.Product{},
		&domain.Store{},
		&domain.UserCart{},
	)
}
