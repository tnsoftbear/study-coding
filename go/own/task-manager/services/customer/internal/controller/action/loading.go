package action

import (
	"fmt"
	"net/http"

	"task_manager/internal/config"
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

func FindParcelLockersByCustomerShippingId(c *gin.Context, rm *storage.RedisManager) {
	id := c.Params.ByName("id")
	shipping, err := rm.LoadCustomerShippingById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
	} else if shipping == nil {
		c.JSON(http.StatusNotFound, gin.H{"id": id, "message": fmt.Sprintf("Customer shipping address not found by id: %s", id)})
	} else {
		// TODO: inject and invert
		plClient := parcel_locker.ParcelLockerClient{
			LocationServiceEndpoint: config.GetStrEnv("LOCATION_HOST", "http://localhost:8081"),
		}
		parcel_lockers, err := plClient.FindParcelLockersNear(shipping, 10000) // TODO: def radius
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"parcel_lockers": parcel_lockers})
	}
}
