package v1

import (
	"context"
	"maktapp-api/internal/apps/api/structs"
	. "maktapp-api/internal/engines/http"
	"net/http"
)

func Entrypoint(ctx context.Context, router *RouteGroup) {
	appInfo := structs.AppInfo{
		Name:    ctx.Value("AppName").(string),
		Version: "1.0.0",
	}

	router.GET("", func(c RequestContext) error {
		return c.JSONPretty(
			http.StatusOK,
			appInfo,
			"  ",
		)
	})
}
