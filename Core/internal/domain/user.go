package domain

type User struct {
	ID     uint32 `json:"id" gorm:"primaryKey"`
	IsAuth bool   `json:"is_auth"`
}

type UserInfo struct {
	Id           uint32 `json:"id" gorm:"primaryKey"`
	Phone        string `json:"phone"`
	Token        string `json:"jwt"`
	RefreshToken string `json:"refresh_jwt"`
}

type UserCart struct {
	UserId    uint32 `json:"user_id"`
	ProductId uint32 `json:"product_id"`
}
