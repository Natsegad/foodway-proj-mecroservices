package router

import (
	"App1/view"
	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	//Users
	router.GET("/", view.ShowTest)
	//router.GET("/users", view.ShowAllUsers)
	//router.POST("/user/add", view.CreateUser)

	router.Run()

}
