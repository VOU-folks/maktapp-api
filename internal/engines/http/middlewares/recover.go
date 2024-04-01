package middlewares

import (
	middlewares "github.com/labstack/echo/v4/middleware"

	"maktapp-api/internal/engines/http"
)

func Recover() http.MiddlewareFunc {
	return middlewares.Recover()
}
