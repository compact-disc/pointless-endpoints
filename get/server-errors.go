package get

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InternalServerError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}

func NotImplemented() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{})
	}
}

func BadGateway() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusBadGateway, gin.H{})
	}
}

func ServiceUnavailable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusServiceUnavailable, gin.H{})
	}
}

func GatewayTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusGatewayTimeout, gin.H{})
	}
}

func HttpVersionNotSupported() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusHTTPVersionNotSupported, gin.H{})
	}
}

func NotExtended() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotExtended, gin.H{})
	}
}

func NetworkAuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{})
	}
}
