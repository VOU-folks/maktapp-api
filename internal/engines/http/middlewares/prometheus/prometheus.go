package prometheus

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"maktapp-api/internal/engines/http"
)

func Middleware(subsystem string) http.MiddlewareFunc {
	return echoprometheus.NewMiddleware(subsystem)
}

func Handler() http.HandlerFunc {
	return echoprometheus.NewHandler()
}
