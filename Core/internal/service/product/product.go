package product

import (
	"core/internal/domain"
	"core/internal/service/store"
	"core/pkg/db"
	"core/pkg/logger"
	"errors"
	"fmt"
)

func addProductToDataBase(prod domain.Product) error {
	return db.DataBase.Create(&prod).Error
}

func CreateProduct(prod domain.Product) error {
	log := logger.GetLogger()
	log.Infof("Created product name %s store name %s ", prod.Name, prod.StoreName)

	err := prod.Validate()
	if err != nil {
		return err
	}

	have, _, err := store.HaveStoreInDb(prod.StoreName)
	if err != nil {
		return err
	}
	if !have {
		return errors.New(fmt.Sprintf("Error: store %s not found ", prod.StoreName))
	}

	return addProductToDataBase(prod)
}
