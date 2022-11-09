package domain

type ResStore struct {
	StoreName string    `json:"store_name"`
	Products  []Product `json:"products"`
}

type Store struct {
	StoreName string `json:"store_name"`
	StoreID   uint32 `json:"store_id"`
}

func (ths *Store) Validate() error {
	return nil
}

type GetStores struct {
	Stores []Store `json:"stores"`
}
