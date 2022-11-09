package product

import (
	"core/internal/domain"
	"core/internal/service/store"
	"core/pkg/db"
	"core/pkg/logger"
	"errors"
	"fmt"
)

func getAllProducts() ([]domain.Product, error) {
	var allProdcuts []domain.Product

	ok := db.DataBase.Find(&allProdcuts)
	return allProdcuts, ok.Error
}

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

	have, id, err := store.HaveStoreInDb(prod.StoreName)
	if err != nil {
		return err
	}
	if !have {
		return errors.New(fmt.Sprintf("Error: store %s not found ", prod.StoreName))
	}

	prod.StoreID = id

	return addProductToDataBase(prod)
}

func GetProductsInStore(id uint32) ([]domain.Product, error) {
	var retProds []domain.Product

	products, err := getAllProducts()
	if err != nil {
		return nil, err
	}

	for _, v := range products {
		if v.StoreID == id {
			retProds = append(retProds, v)
		}
	}

	return retProds, nil
}
