package route

import (
	"task_manager/internal/controller/action"
	"task_manager/internal/config"
	"task_manager/internal/parcel_locker"
	"task_manager/internal/storage"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	rm := storage.NewRedisManager()
	plClient := &parcel_locker.ParcelLockerClient{
		LocationServiceEndpoint: config.GetStrEnv("PARCEL_LOCKER_SERVICE_ADDR", "http://localhost:8081"),
	}

	r.GET("/ping", func(c *gin.Context) { action.Ping(c) })
	r.GET("/customer/shippings", func(c *gin.Context) { action.LoadAllCustomerShippings(c, rm) })
	r.GET("/customer/shipping/:id", func(c *gin.Context) { action.LoadAllCustomerShippings(c, rm) })
	r.GET("/customer/shipping-parcel-lockers/:id", func(c *gin.Context) { action.FindParcelLockersByCustomerShippingId(c, rm, plClient) })
	r.POST("/customer/shipping", func(c *gin.Context) { action.SaveCustomerShipping(c, rm) })
	r.DELETE("/customer/shipping/:id", func(c *gin.Context) { action.DeleteCustomerShipping(c, rm) })
	return r
}
