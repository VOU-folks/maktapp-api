package middlewares

import (
	middlewares "github.com/labstack/echo/v4/middleware"

	"maktapp-api/internal/engines/http"
)

func RequestId() http.MiddlewareFunc {
	return middlewares.RequestID()
}
