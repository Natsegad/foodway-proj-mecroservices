package service

import (
	"core/internal/domain"
	"core/pkg/logger"
)

func addProductToDataBase(prod domain.Product) {

}

func CreateProduct(prod domain.Product) error {
	log := logger.GetLogger()
	log.Infof("Created product name %s store name %s ", prod.Name, prod.StoreName)

	err := prod.Validate()
	if err != nil {
		return err
	}

	addProductToDataBase(prod)

	return nil
}
