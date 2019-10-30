package servers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReportFormat ...
func ReportFormat(c *gin.Context, success bool, msg string, json gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"msg":     msg,
		"data":    json,
	})
}
