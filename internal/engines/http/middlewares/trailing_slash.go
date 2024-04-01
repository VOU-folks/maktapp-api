package middlewares

import (
	middlewares "github.com/labstack/echo/v4/middleware"
	statusCodes "net/http"

	"maktapp-api/internal/engines/http"
)

func AddTrailingSlash() http.MiddlewareFunc {
	return middlewares.AddTrailingSlashWithConfig(
		middlewares.TrailingSlashConfig{
			RedirectCode: statusCodes.StatusMovedPermanently,
		})
}

func RemoveTrailingSlash() http.MiddlewareFunc {
	return middlewares.RemoveTrailingSlashWithConfig(
		middlewares.TrailingSlashConfig{
			RedirectCode: statusCodes.StatusMovedPermanently,
		})
}
