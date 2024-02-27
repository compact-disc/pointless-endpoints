package main

import (
	"github.com/gin-gonic/gin"
	"pointless-endpoints/get"
)

func main() {
	r := gin.Default()

	/*
		Get Endpoint Handlers
	*/
	// Informational - 100
	r.GET("/continue", get.Continue())

	// Success - 200
	r.GET("/ok", get.Ok())
	r.GET("/created", get.Created())
	r.GET("/accepted", get.Accepted())
	r.GET("/noContent", get.NoContent())

	// Redirection - 300
	r.GET("/movedPermanently", get.MovedPermanently())
	r.GET("/found", get.Found())
	r.GET("/temporaryRedirect", get.TemporaryRedirect())
	r.GET("/permanentRedirect", get.PermanentRedirect())

	// Client Errors - 400
	r.GET("/badRequest", get.BadRequest())
	r.GET("/unauthorized", get.Unauthorized())
	r.GET("/forbidden", get.Forbidden())
	r.GET("/notFound", get.NotFound())
	r.GET("/methodNotAllowed", get.MethodNotAllowed())
	r.GET("/notAcceptable", get.NotAcceptable())
	r.GET("/proxyAuthRequired", get.ProxyAuthRequired())
	r.GET("/requestTimeout", get.RequestTimeout())
	r.GET("/conflict", get.Conflict())
	r.GET("/gone", get.Conflict())
	r.GET("/lengthRequired", get.LengthRequired())
	r.GET("/preconditionFailed", get.PreconditionFailed())
	r.GET("/payloadTooLarge", get.PayloadTooLarge())
	r.GET("/uriTooLong", get.UriTooLong())
	r.GET("/unsupportedMediaType", get.UnsupportedMediaType())
	r.GET("/rangeNotSatisfiable", get.RangeNotSatisfiable())
	r.GET("/expectationFailed", get.ExpectationFailed())
	r.GET("/iAmTeapot", get.ImATeapot())
	r.GET("/misdirectedRequest", get.MisdirectedRequest())
	r.GET("/unprocessableContent", get.UnprocessableContent())
	r.GET("/tooManyRequests", get.TooManyRequests())
	r.GET("/unavailableForLegalReasons", get.UnavailableForLegalReasons())

	// Server Errors - 500
	r.GET("/internalServerError", get.InternalServerError())
	r.GET("/notImplemented", get.NotImplemented())
	r.GET("/badGateway", get.BadGateway())
	r.GET("/serviceUnavailable", get.ServiceUnavailable())
	r.GET("/gatewayTimeout", get.GatewayTimeout())
	r.GET("/httpVersionNotSupported", get.HttpVersionNotSupported())
	r.GET("/notExtended", get.NotExtended())
	r.GET("/networkAuthenticationRequired", get.NetworkAuthenticationRequired())

	r.Run()
}
