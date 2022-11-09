package store

import "core/internal/domain"

func addStoreToDb(store domain.Store) error {
	return nil
}

func CreateStore(store domain.Store) error {
	err := store.Validate()
	if err != nil {
		return err
	}

	return addStoreToDb(store)
}
