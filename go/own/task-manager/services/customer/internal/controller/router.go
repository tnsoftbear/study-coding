package controller

import (
	"fmt"
	"net/http"

	"task_manager/internal/config"
	"task_manager/internal/location"
	"task_manager/internal/redis_storage"
	"task_manager/internal/types"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	rm := redis_storage.NewRedisManager()

	r.GET("/customer/shippings", func(c *gin.Context) {
		if shippings, err := rm.LoadAllCustomerShippings(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			response := gin.H{"shippings": shippings}
			c.JSON(http.StatusOK, response)
		}
	})

	r.GET("/customer/shipping/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		shipping, err := rm.LoadCustomerShippingById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
		} else if shipping == nil {
			c.JSON(http.StatusNotFound, gin.H{"id": id, "message": fmt.Sprintf("Customer shipping address not found by id: %s", id)})
		} else {
			c.JSON(http.StatusOK, gin.H{"shippings": shipping})
		}
	})

	r.GET("/customer/shipping-parcel-lockers/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		shipping, err := rm.LoadCustomerShippingById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
		} else if shipping == nil {
			c.JSON(http.StatusNotFound, gin.H{"id": id, "message": fmt.Sprintf("Customer shipping address not found by id: %s", id)})
		} else {
			// TODO: inject and invert
			locationService := location.LocationService{
				LocationServiceEndpoint: config.GetStrEnv("LOCATION_HOST", "http://localhost:8081"),
			}
			parcel_lockers, err := locationService.FindParcelLockersNear(shipping, 10000) // TODO: def radius
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"id": id, "message": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"parcel_lockers": parcel_lockers})
		}
	})

	r.POST("/customer/shipping", func(c *gin.Context) {
		var shipping types.CustomerShipping = types.CustomerShipping{}
		if err := c.BindJSON(&shipping); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"shipping": shipping, "created": false, "message": err.Error()})
			return
		}

		if err := rm.SaveCustomerShipping(shipping); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"shipping": shipping, "created": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"shipping": shipping, "created": true, "message": "Customer shipping address created successfully"})
	})

	r.DELETE("/customer/shipping/:id", func(c *gin.Context) {
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
	})

	return r
}
