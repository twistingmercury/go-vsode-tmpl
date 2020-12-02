package handler

import (
	"github.com/gin-gonic/gin"
)

// Get ..
func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Hello, authenticated client!")
	}
}
