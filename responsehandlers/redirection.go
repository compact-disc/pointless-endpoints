package responsehandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MovedPermanently() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMovedPermanently, gin.H{})
	}
}

func Found() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusFound, gin.H{})
	}
}

func TemporaryRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusTemporaryRedirect, gin.H{})
	}
}

func PermanentRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusPermanentRedirect, gin.H{})
	}
}
