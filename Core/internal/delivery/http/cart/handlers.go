package cart

import (
	"core/internal/domain"
	"core/internal/service/user"
	"core/pkg/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

func GetCart(c *gin.Context) {
	log := logger.GetLogger()

	qId := c.Request.URL.Query().Get("id")

	uId, err := strconv.ParseUint(qId, 10, 32)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	products, err := user.GetUserCart(uint32(uId))
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, products)
}

func AddProductInCart(c *gin.Context) {
	pInCartS := domain.UserCart{}

	log := logger.GetLogger()

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #1",
		})
		return
	}

	err = json.Unmarshal(data, &pInCartS)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #2",
		})
		return
	}

	err = user.AddProductInCart(pInCartS)
	if err != nil {
		log.Errorf("Error read Request Body %s ", err.Error())
		c.JSON(400, gin.H{
			"error": "Error bad request #2",
		})
		return
	}

	c.JSON(200, "ok")
}
