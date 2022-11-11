package user

import (
	"core/internal/domain"
	"core/pkg/db"
)

func getUserCarts() ([]domain.UserCart, error) {
	var carts []domain.UserCart

	err := db.DataBase.Find(&carts).Error

	return carts, err
}

func getUserInfoById(uId uint32) (ret domain.UserCart, err error) {
	err = db.DataBase.Find(&ret, uId).Error
	return
}

func GetUserCart(uId uint32) ([]uint32, error) {
	var cartProducts []uint32

	carts, err := getUserCarts()
	if err != nil {
		return nil, err
	}

	for _, v := range carts {
		if v.UserId == uId {
			cartProducts = append(cartProducts, v.ProductId)
		}
	}

	return cartProducts, nil
}

func addProductInCartDb(info domain.UserCart) error {
	return db.DataBase.Create(&info).Error
}

func AddProductInCart(info domain.UserCart) error {
	return addProductInCartDb(info)
}
