package domain

type Store struct {
	StoreName string `json:"store_name"`
	StoreID   uint32 `json:"store_id"`
}

func (ths *Store) Validate() error {
	return nil
}
