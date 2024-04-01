package api

import (
	"context"

	v1 "maktapp-api/internal/apps/api/handlers/v1"
	. "maktapp-api/internal/engines/http"
)

var routeRegistrarMap map[string]RouteRegistrar = map[string]RouteRegistrar{
	"/v1": v1.Entrypoint,
}

func RegisterRoutes(context context.Context, engine *RouteEngine) {
	for route, registrar := range routeRegistrarMap {
		router := engine.Group(route)
		registrar(context, router)
	}
}
