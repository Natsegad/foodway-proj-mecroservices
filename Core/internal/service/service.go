package service

import (
	"core/internal/domain"
	"core/pkg/logger"
)

func CreateProduct(prod domain.Product) error {
	log := logger.GetLogger()
	log.Infof("Created product name %s store name %s ", prod.Name, prod.StoreName)


	return nil
}