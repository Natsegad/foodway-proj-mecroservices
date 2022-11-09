package store

import (
	"core/internal/domain"
	"core/internal/service/store"
	"core/pkg/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
)

func CreateStore(c *gin.Context) {
	log := logger.GetLogger()

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #1",
		})
		return
	}

	storeCreateReq := domain.StoreCreateReq{}

	err = json.Unmarshal(data, &storeCreateReq)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #2",
		})
		return
	}

	storeId := uuid.New().ID()
	storeNew := domain.Store{
		StoreID:   storeId,
		StoreName: storeCreateReq.StoreName,
	}

	err = store.CreateStore(storeNew)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #2",
		})
		return
	}

	c.JSON(200, storeNew)
}

func GetProductInStore(c *gin.Context) {
	
}

func GetStores(c *gin.Context) {

}
