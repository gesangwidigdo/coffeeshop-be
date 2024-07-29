package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindData(obj interface{}, c *gin.Context) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		ReturnResponse(http.StatusBadRequest, "Failed to Bind Data", "error", err.Error(), c)
		return false
	}
	return true
}
