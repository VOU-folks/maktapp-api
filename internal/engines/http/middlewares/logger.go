package middlewares

import (
	middlewares "github.com/labstack/echo/v4/middleware"

	"maktapp-api/internal/engines/http"
)

func Logger() http.MiddlewareFunc {
	return middlewares.LoggerWithConfig(
		middlewares.LoggerConfig{
			Format: `${time_rfc3339} | ${remote_ip} | ${method} ${uri} | ${status} | ${latency_human} | ${bytes_in} / ${bytes_out} (in/out) | ${id}` + "\n",
		},
	)
}
