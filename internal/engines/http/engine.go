package http

import (
	"context"
	"github.com/labstack/echo/v4"
)

type RouteEngine = echo.Echo
type RouteGroup = echo.Group
type RouteRegistrar = func(context.Context, *RouteGroup)
type RequestContext = echo.Context
type HandlerFunc = echo.HandlerFunc
type MiddlewareFunc = echo.MiddlewareFunc

func NewRouteEngine() *RouteEngine {
	return echo.New()
}
