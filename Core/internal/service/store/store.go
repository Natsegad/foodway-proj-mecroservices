package store

import (
	"core/internal/domain"
	"core/pkg/db"
	"errors"
	"fmt"
)

func getAllStores() ([]domain.Store, error) {
	var stores []domain.Store

	ok := db.DataBase.Find(&stores)
	if ok.Error != nil {
		return nil, ok.Error
	}

	return stores, nil
}

func haveStoreInDb(storeName string) (bool, uint32, error) {
	stores, err := getAllStores()
	if err != nil {
		return false, 0, err
	}

	for _, v := range stores {
		if v.StoreName == storeName {
			return true, v.StoreID, nil
		}
	}

	return false, 0, nil
}

func addStoreToDb(store domain.Store) error {
	return db.DataBase.Create(&store).Error
}

func CreateStore(store domain.Store) error {
	err := store.Validate()
	if err != nil {
		return err
	}

	have, _, err := haveStoreInDb(store.StoreName)
	if err != nil {
		return err
	}
	if have {
		return errors.New(fmt.Sprintf("store %s have", store.StoreName))
	}

	return addStoreToDb(store)
}

func GetStores() ([]domain.Store, error) {
	return getAllStores()
}
