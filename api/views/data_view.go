package views

import "github.com/gin-gonic/gin"

func DataView(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"error":   false,
		"message": message,
		"data":    data,
	})
}
