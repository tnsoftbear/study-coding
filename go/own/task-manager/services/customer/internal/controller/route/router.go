package route

import (
	"task_manager/internal/controller/action"
	"task_manager/internal/storage"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	rm := storage.NewRedisManager()
	r.GET("/ping", func(c *gin.Context) { action.Ping(c) })
	r.GET("/customer/shippings", func(c *gin.Context) { action.LoadAllCustomerShippings(c, rm) })
	r.GET("/customer/shipping/:id", func(c *gin.Context) { action.LoadAllCustomerShippings(c, rm) })
	r.GET("/customer/shipping-parcel-lockers/:id", func(c *gin.Context) { action.FindParcelLockersByCustomerShippingId(c, rm) })
	r.POST("/customer/shipping", func(c *gin.Context) { action.SaveCustomerShipping(c, rm) })
	r.DELETE("/customer/shipping/:id", func(c *gin.Context) { action.DeleteCustomerShipping(c, rm) })
	return r
}
