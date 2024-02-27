package get

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
}

func Unauthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
}

func Forbidden() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusForbidden, gin.H{})
	}
}

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}

func MethodNotAllowed() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{})
	}
}

func NotAcceptable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotAcceptable, gin.H{})
	}
}

func ProxyAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusProxyAuthRequired, gin.H{})
	}
}

func RequestTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusRequestTimeout, gin.H{})
	}
}

func Conflict() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusConflict, gin.H{})
	}
}

func Gone() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusGone, gin.H{})
	}
}

func LengthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusLengthRequired, gin.H{})
	}
}

func PreconditionFailed() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusPreconditionFailed, gin.H{})
	}
}

func PayloadTooLarge() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{})
	}
}

func UriTooLong() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{})
	}
}

func UnsupportedMediaType() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{})
	}
}

func RangeNotSatisfiable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{})
	}
}

func ExpectationFailed() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusExpectationFailed, gin.H{})
	}
}

func ImATeapot() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusTeapot, gin.H{})
	}
}

func MisdirectedRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMisdirectedRequest, gin.H{})
	}
}

func UnprocessableContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{})
	}
}

func TooManyRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusTooManyRequests, gin.H{})
	}
}

func UnavailableForLegalReasons() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusUnavailableForLegalReasons, gin.H{})
	}
}
