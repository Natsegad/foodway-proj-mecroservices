package product

import (
	"core/internal/domain"
	"core/internal/service/product"
	"core/pkg/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"strconv"
)

func CreateProductH(c *gin.Context) {
	log := logger.GetLogger()

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #1",
		})
		return
	}

	prodResp := domain.Product{}
	prodResp.ID = uuid.New().ID()

	err = json.Unmarshal(data, &prodResp)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #2",
		})
		return
	}

	err = product.CreateProduct(prodResp)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, prodResp.StoreName)
}

func GetProductInStore(c *gin.Context) {
	log := logger.GetLogger()

	storeIdS := c.Request.URL.Query().Get("id")
	uId, err := strconv.ParseUint(storeIdS, 10, 32)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	products, err := product.GetProductsInStore(uint32(uId))
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #2",
		})
		return
	}

	c.JSON(200, products)
}
