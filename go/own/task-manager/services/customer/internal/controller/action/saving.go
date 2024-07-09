package action

import (
	"net/http"

	"task_manager/internal/model"
	"task_manager/internal/storage"

	"github.com/gin-gonic/gin"
)

func SaveCustomerShipping(c *gin.Context, rm *storage.RedisManager) {
	var shipping model.CustomerShipping = model.CustomerShipping{}
	if err := c.BindJSON(&shipping); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"shipping": shipping, "created": false, "message": err.Error()})
		return
	}

	if err := rm.SaveCustomerShipping(shipping); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{"shipping": shipping, "created": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"shipping": shipping, "created": true, "message": "Customer shipping address created successfully"})
}
