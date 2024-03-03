package main

import (
	"github.com/gin-gonic/gin"
	"pointless-endpoints/responsehandlers"
)

const (
	GET    int = 0
	PUT        = 1
	POST       = 2
	PATCH      = 3
	DELETE     = 4
)

func main() {
	router := gin.Default()

	addHandlers(router)

	router.Run()
}

func addHandlers(router *gin.Engine) {

	for i := 0; i <= DELETE; i++ {

		// Informational - 100
		addHandler(i, router, "/continue", responsehandlers.Continue())

		// Success - 200
		addHandler(i, router, "/ok", responsehandlers.Ok())
		addHandler(i, router, "/created", responsehandlers.Created())
		addHandler(i, router, "/accepted", responsehandlers.Accepted())
		addHandler(i, router, "/noContent", responsehandlers.NoContent())

		// Redirection - 300
		addHandler(i, router, "/movedPermanently", responsehandlers.MovedPermanently())
		addHandler(i, router, "/found", responsehandlers.Found())
		addHandler(i, router, "/temporaryRedirect", responsehandlers.TemporaryRedirect())
		addHandler(i, router, "/permanentRedirect", responsehandlers.PermanentRedirect())

		// Client Errors - 400
		addHandler(i, router, "/badRequest", responsehandlers.BadRequest())
		addHandler(i, router, "/unauthorized", responsehandlers.Unauthorized())
		addHandler(i, router, "/forbidden", responsehandlers.Forbidden())
		addHandler(i, router, "/notFound", responsehandlers.NotFound())
		addHandler(i, router, "/methodNotAllowed", responsehandlers.MethodNotAllowed())
		addHandler(i, router, "/notAcceptable", responsehandlers.NotAcceptable())
		addHandler(i, router, "/proxyAuthRequired", responsehandlers.ProxyAuthRequired())
		addHandler(i, router, "/requestTimeout", responsehandlers.RequestTimeout())
		addHandler(i, router, "/conflict", responsehandlers.Conflict())
		addHandler(i, router, "/gone", responsehandlers.Gone())
		addHandler(i, router, "/lengthRequired", responsehandlers.LengthRequired())
		addHandler(i, router, "/preconditionFailed", responsehandlers.PreconditionFailed())
		addHandler(i, router, "/payloadTooLarge", responsehandlers.PayloadTooLarge())
		addHandler(i, router, "/uriTooLong", responsehandlers.UriTooLong())
		addHandler(i, router, "/unsupportedMediaType", responsehandlers.UnsupportedMediaType())
		addHandler(i, router, "/rangeNotSatisfiable", responsehandlers.RangeNotSatisfiable())
		addHandler(i, router, "/expectationFailed", responsehandlers.ExpectationFailed())
		addHandler(i, router, "/iAmTeapot", responsehandlers.ImATeapot())
		addHandler(i, router, "/misdirectedRequest", responsehandlers.MisdirectedRequest())
		addHandler(i, router, "/unprocessableContent", responsehandlers.UnprocessableContent())
		addHandler(i, router, "/tooManyRequests", responsehandlers.TooManyRequests())
		addHandler(i, router, "/unavailableForLegalReasons", responsehandlers.UnavailableForLegalReasons())

		// Server Errors - 500
		addHandler(i, router, "/internalServerError", responsehandlers.InternalServerError())
		addHandler(i, router, "/notImplemented", responsehandlers.NotImplemented())
		addHandler(i, router, "/badGateway", responsehandlers.BadGateway())
		addHandler(i, router, "/serviceUnavailable", responsehandlers.ServiceUnavailable())
		addHandler(i, router, "/gatewayTimeout", responsehandlers.GatewayTimeout())
		addHandler(i, router, "/httpVersionNotSupported", responsehandlers.HttpVersionNotSupported())
		addHandler(i, router, "/notExtended", responsehandlers.NotExtended())
		addHandler(i, router, "/networkAuthenticationRequired", responsehandlers.NetworkAuthenticationRequired())
	}
}

func addHandler(
	method int,
	router *gin.Engine,
	relativePath string,
	handlerFunc gin.HandlerFunc,
) {
	switch method {
	case GET:
		router.GET(relativePath, handlerFunc)
		break
	case PUT:
		router.PUT(relativePath, handlerFunc)
		break
	case POST:
		router.POST(relativePath, handlerFunc)
		break
	case PATCH:
		router.PATCH(relativePath, handlerFunc)
		break
	case DELETE:
		router.DELETE(relativePath, handlerFunc)
		break
	}
}
