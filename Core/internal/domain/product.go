package domain

import (
	"core/pkg/db"
	"fmt"
)

type Product struct {
	ID        uint32 `json:"ID" gorm:"primaryKey"`
	StoreID   uint32 `json:"store_id"`
	StoreName string `json:"store_name"`
	Name      string `json:"name"`
	Price     uint32 `json:"price"`
	ImagePath string `json:"image_path"`
	Grams     uint32 `json:"grams"`
	Calories  uint32 `json:"calories"`
	Sqrls     uint32 `json:"sqrls"` // белки
	Fats      uint32 `json:"fats"`
	Carbonts  uint32 `json:"carbonts"` // углеводы
}

type ProductEmbeds struct {
	P         Product `json:"product"`
	ImagePath string  `json:"image_path"`
	U         User    `json:"user"`
}

func (ths *Product) Validate() error {

	return nil
}

func (ths *Product) AddToDataBase() error {
	ok := db.DataBase.Create(&(*ths))
	if ok.Error != nil {
		fmt.Printf("Error Product %s \n", ok.Error.Error())
		return ok.Error
	}

	fmt.Printf("Product added to database %s store name %s \n", ths.Name, ths.StoreName)
	return nil
}

func (ths *Product) DeleteProductDataBase() error {
	user := Product{}
	ok := db.DataBase.Find(&user, ths.ID)
	if ok.Error != nil {
		fmt.Printf("Error Delete Product %s \n", ok.Error.Error())
		return ok.Error
	}

	db.DataBase.Delete(&user)

	return nil
}

func GetProductById(id uint32) (Product, error) {
	user := Product{}
	ok := db.DataBase.Find(&user, id)
	if ok.Error != nil {
		fmt.Printf("Error GetProductById %s \n", ok.Error.Error())
		return user, ok.Error
	}

	return user, nil
}

func NewProduct(id, userId, price, grams, calories, sqrls, fats, carbonts uint32, storeName, name, imagePath string) Product {
	return Product{
		ID:        id,
		StoreID:   userId,
		StoreName: storeName,
		Name:      name,
		Price:     price,
		ImagePath: imagePath,
		Grams:     grams,
		Calories:  calories,
		Sqrls:     sqrls,
		Fats:      fats,
		Carbonts:  carbonts,
	}
}
