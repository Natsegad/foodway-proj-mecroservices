package app

import (
	cfg2 "core/internal/cfg"
	"core/internal/delivery/http/product"
	"core/internal/service/db"
	pkgdb "core/pkg/db"
	"github.com/gin-gonic/gin"
)

func Start() {
	cfg2.LoadEnv()
	cfg2.InitCfg()
	pkgdb.InitDb()
	db.AutoMigrateService()

	router := gin.Default()

	prod := router.Group("products/")
	{
		prod.GET("/", func(context *gin.Context) {
			context.JSON(200, "main page product")
		})
		prod.POST("/create", product.CreateProductH)
	}

	router.Run(cfg2.Cfg.IP + ":" + cfg2.Cfg.Port)
}
