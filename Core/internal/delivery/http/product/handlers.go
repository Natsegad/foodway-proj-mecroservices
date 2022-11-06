package product

import (
	"core/internal/domain"
	"core/internal/service"
	"core/pkg/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
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

	err = json.Unmarshal(data, &prodResp)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #2",
		})
		return
	}

	err = service.CreateProduct(prodResp)

	c.JSON(200, prodResp)
}
