package utils

import "github.com/gin-gonic/gin"

func ReturnResponse(StatusCode int, Message string, Key string, data interface{}, c *gin.Context) {
	if (Key == "" && data == nil) {
		c.JSON(StatusCode, gin.H{
			"message": Message,
		})
	} else {
		c.JSON(StatusCode, gin.H{
			"message": Message,
			Key: data,
		})
	}
}