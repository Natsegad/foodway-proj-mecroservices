package app

import (
	cfg2 "core/internal/cfg"
	"core/internal/delivery/http/cart"
	"core/internal/delivery/http/product"
	store2 "core/internal/delivery/http/store"
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
	prod.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
	})
	{
		prod.GET("/", func(context *gin.Context) {
			context.JSON(200, "main page product")
		})
		prod.POST("create/", product.CreateProductH)
		prod.GET("getProduct/", product.GetProductInStore)
		prod.GET("getProducts/", product.GetAllProducts)
	}

	store := router.Group("store/")
	store.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
	})
	{
		store.GET("/", func(context *gin.Context) {
			context.JSON(200, "main page store")
		})
		store.POST("create/", store2.CreateStore)
		store.GET("getStores/", store2.GetStores)
	}

	cartPage := router.Group("cart/")
	cartPage.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
	})
	{
		cartPage.GET("/", func(context *gin.Context) {
			context.JSON(200, "main page cart")
		})
		cartPage.POST("addProductInCart/", cart.AddProductInCart)
		cartPage.GET("getProductsInCart/", cart.GetCart)
	}

	router.Run(cfg2.Cfg.IP + ":" + cfg2.Cfg.Port)
}
