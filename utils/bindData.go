package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

func BindData(obj interface{}, c *gin.Context) bool {
	g := galidator.New()
	customizer := g.Validator(obj)

	if err := c.ShouldBindJSON(obj); err != nil {
		ReturnResponse(http.StatusBadRequest, "Failed to Bind Data", "error", customizer.DecryptErrors(err), c)
		return false
	}
	return true
}
