package domain

type User struct {
	ID     uint32 `json:"id" gorm:"primaryKey"`
	IsAuth bool   `json:"is_auth"`
}
