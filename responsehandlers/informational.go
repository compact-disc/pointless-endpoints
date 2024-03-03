package responsehandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Continue() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusContinue, gin.H{})
	}
}
