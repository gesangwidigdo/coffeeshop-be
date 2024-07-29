package main

import (
	"net/http"

	"github.com/gesangwidigdo/coffeeshop-be/initializers"
	"github.com/gesangwidigdo/coffeeshop-be/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "oh, hi",
		})
	})

	// Outlet Route
	outletRoutes := r.Group("/outlet")
	routes.OutletRoute(outletRoutes)

	// Position Route
	positionRoutes := r.Group("/position")
	routes.PositionRoute(positionRoutes)
	
	r.Run()
}
