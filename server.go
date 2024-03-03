package main

import (
	"github.com/gin-gonic/gin"
	"pointless-endpoints/responsehandlers"
)

func main() {
	router := gin.Default()

	addGetHandlers(router)

	router.Run()
}

func addGetHandlers(router *gin.Engine) {
	// Informational - 100
	router.GET("/continue", responsehandlers.Continue())

	// Success - 200
	router.GET("/ok", responsehandlers.Ok())
	router.GET("/created", responsehandlers.Created())
	router.GET("/accepted", responsehandlers.Accepted())
	router.GET("/noContent", responsehandlers.NoContent())

	// Redirection - 300
	router.GET("/movedPermanently", responsehandlers.MovedPermanently())
	router.GET("/found", responsehandlers.Found())
	router.GET("/temporaryRedirect", responsehandlers.TemporaryRedirect())
	router.GET("/permanentRedirect", responsehandlers.PermanentRedirect())

	// Client Errors - 400
	router.GET("/badRequest", responsehandlers.BadRequest())
	router.GET("/unauthorized", responsehandlers.Unauthorized())
	router.GET("/forbidden", responsehandlers.Forbidden())
	router.GET("/notFound", responsehandlers.NotFound())
	router.GET("/methodNotAllowed", responsehandlers.MethodNotAllowed())
	router.GET("/notAcceptable", responsehandlers.NotAcceptable())
	router.GET("/proxyAuthRequired", responsehandlers.ProxyAuthRequired())
	router.GET("/requestTimeout", responsehandlers.RequestTimeout())
	router.GET("/conflict", responsehandlers.Conflict())
	router.GET("/gone", responsehandlers.Gone())
	router.GET("/lengthRequired", responsehandlers.LengthRequired())
	router.GET("/preconditionFailed", responsehandlers.PreconditionFailed())
	router.GET("/payloadTooLarge", responsehandlers.PayloadTooLarge())
	router.GET("/uriTooLong", responsehandlers.UriTooLong())
	router.GET("/unsupportedMediaType", responsehandlers.UnsupportedMediaType())
	router.GET("/rangeNotSatisfiable", responsehandlers.RangeNotSatisfiable())
	router.GET("/expectationFailed", responsehandlers.ExpectationFailed())
	router.GET("/iAmTeapot", responsehandlers.ImATeapot())
	router.GET("/misdirectedRequest", responsehandlers.MisdirectedRequest())
	router.GET("/unprocessableContent", responsehandlers.UnprocessableContent())
	router.GET("/tooManyRequests", responsehandlers.TooManyRequests())
	router.GET("/unavailableForLegalReasons", responsehandlers.UnavailableForLegalReasons())

	// Server Errors - 500
	router.GET("/internalServerError", responsehandlers.InternalServerError())
	router.GET("/notImplemented", responsehandlers.NotImplemented())
	router.GET("/badGateway", responsehandlers.BadGateway())
	router.GET("/serviceUnavailable", responsehandlers.ServiceUnavailable())
	router.GET("/gatewayTimeout", responsehandlers.GatewayTimeout())
	router.GET("/httpVersionNotSupported", responsehandlers.HttpVersionNotSupported())
	router.GET("/notExtended", responsehandlers.NotExtended())
	router.GET("/networkAuthenticationRequired", responsehandlers.NetworkAuthenticationRequired())
}
