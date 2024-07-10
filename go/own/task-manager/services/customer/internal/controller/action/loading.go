package action

import (
	"fmt"
	"net/http"
	"strconv"

	"task_manager/internal/parcel_locker"
	"task_manager/internal/storage"

	"github.com/gin-gonic/gin"
)

func LoadAllCustomerShippings(c *gin.Context, rm *storage.RedisManager) {
	if shippings, err := rm.LoadAllCustomerShippings(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		response := gin.H{"shippings": shippings}
		c.JSON(http.StatusOK, response)
	}
}

func LoadCustomerShippingById(c *gin.Context, rm *storage.RedisManager) {
	id := c.Params.ByName("id")
	shipping, err := rm.LoadCustomerShippingById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
	} else if shipping == nil {
		c.JSON(http.StatusNotFound, gin.H{"id": id, "message": fmt.Sprintf("Customer shipping address not found by id: %s", id)})
	} else {
		c.JSON(http.StatusOK, gin.H{"shipping": shipping})
	}
}

func FindParcelLockersByCustomerShippingId(
	c *gin.Context, 
	rm *storage.RedisManager, 
	plClient *parcel_locker.ParcelLockerClient,
) {
	id := c.Params.ByName("id")
	distanceStr := c.DefaultQuery("distance", "10")
	distance, err := strconv.ParseFloat(distanceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid distance value"})
		// c.Abort()
		return
	}

	shipping, err := rm.LoadCustomerShippingById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
	} else if shipping == nil {
		c.JSON(http.StatusNotFound, gin.H{"id": id, "message": fmt.Sprintf("Customer shipping address not found by id: %s", id)})
	} else {
		parcel_lockers, err := plClient.FindParcelLockersNear(shipping, distance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"parcel_lockers": parcel_lockers})
	}
}
