package view

import (
	"App1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	//{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	//{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func ShowTest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func ShowUserInfo(context *gin.Context) models.User {
	var user models.User
	return user
}

//func ShowAllUsers(context *gin.Context) {
//	var users []models.User
//
//	models.DB().Find(&users)
//
//	context.JSON(http.StatusOK, gin.H{"users": users})
//}

type CreateUserInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImgSrc      string `json:"img-src" binding:"required"`
	Time        string `json:"time" binding:"required"`
	Read        string `json:"read" binding:"required"`
	Url         string `json:"url" binding:"required"`
}

func CreateUser(context *gin.Context) {
	var userInput CreateUserInput

	if err := context.ShouldBindJSON(&userInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + " Неправильно введены данные"})
		return
	}

	user := models.User{
		Name:        userInput.Name,
		Description: userInput.Description,
		ImgSrc:      userInput.ImgSrc,
		Time:        userInput.Time,
		Read:        userInput.Read,
		Url:         userInput.Url,
	}

	//models.DB().Create(&user)

	context.JSON(http.StatusOK, gin.H{"user": user})

}
