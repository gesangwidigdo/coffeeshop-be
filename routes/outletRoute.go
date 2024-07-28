package routes

import (
	"github.com/gesangwidigdo/coffeeshop-be/controllers"
	"github.com/gin-gonic/gin"
)

func OutletRoute(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllOutlet)
	r.GET("/:id", controllers.GetOutletByID)
	r.POST("/", controllers.CreateOutlet)
	r.PUT("/:id", controllers.UpdateOutlet)
	r.DELETE("/:id", controllers.DeleteOutlet)
}
