package routes

import (
	"github.com/gesangwidigdo/coffeeshop-be/controllers"
	"github.com/gin-gonic/gin"
)

func PositionRoute(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllPosition)
	r.GET("/:id", controllers.GetPositionByID)
	r.POST("/", controllers.CreatePosition)
	r.PUT("/:id", controllers.UpdatePosition)
	r.DELETE("/:id", controllers.DeletePosition)
}
