package routes

import (
	"github.com/gesangwidigdo/coffeeshop-be/controllers"
	"github.com/gin-gonic/gin"
)

func IngredientRoute(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateIngredient)
	r.GET("/", controllers.GetAllIngredient)
	r.GET("/:id", controllers.GetIngredientByID)
	r.PUT("/:id", controllers.UpdateIngredient)
	r.PUT("/stock/:id", controllers.AddStock)
	r.DELETE("/:id", controllers.DeleteIngredient)
}
