package action

import (
	"fmt"
	"net/http"

	"task_manager/internal/storage"

	"github.com/gin-gonic/gin"
)

func DeleteCustomerShipping(c *gin.Context, rm *storage.RedisManager) {
	id := c.Params.ByName("id")
	if shipping, _ := rm.LoadCustomerShippingById(id); shipping == nil {
		c.JSON(http.StatusNotFound, gin.H{"id": id, "deleted": false, "message": fmt.Sprintf("Customer shipping address not found by id: %s", id)})
		return
	}

	if err := rm.DeleteCustomerShipping(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "deleted": true, "message": "Customer shipping address successfully deleted"})
}
