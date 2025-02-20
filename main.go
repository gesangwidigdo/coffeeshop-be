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

	routes.AllRoute(r)
	
	r.Run()
}
