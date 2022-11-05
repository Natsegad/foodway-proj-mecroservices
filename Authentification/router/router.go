package router

import (
	"Authentification/view"
	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	//auth microservice
	router.POST("/auth", view.RedirectAuth)
	router.GET("/auth", view.RedirectAuth)
	router.DELETE("/auth", view.RedirectAuth)

	router.Run()

}
