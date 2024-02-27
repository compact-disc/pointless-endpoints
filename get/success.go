package get

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}

func Created() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{})
	}
}

func Accepted() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{})
	}
}

func NoContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}
