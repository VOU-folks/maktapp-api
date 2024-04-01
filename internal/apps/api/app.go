package api

import (
	"context"
	"errors"
	"maktapp-api/internal/engines/http/middlewares"
	"maktapp-api/internal/engines/http/middlewares/prometheus"
	"net"
	"net/http"

	"golang.org/x/crypto/acme/autocert"

	"maktapp-api/internal/apps/api/structs"
	"maktapp-api/internal/di"
	. "maktapp-api/internal/engines/http"
)

type App struct {
	ctx    context.Context
	di     di.Container
	engine *RouteEngine
}

func NewApp(ctx context.Context, di di.Container) *App {
	appInfo := structs.AppInfo{
		Name:    ctx.Value("AppName").(string),
		Version: ctx.Value("AppVersion").(string),
	}

	engine := NewRouteEngine()
	engine.HideBanner = true

	sslEnabled := ctx.Value("SslEnabled").(bool)
	if sslEnabled {
		engine.AutoTLSManager.HostPolicy = autocert.HostWhitelist(ctx.Value("TlsManagerHostWhitelist").(string))
		engine.AutoTLSManager.Cache = autocert.DirCache(ctx.Value("TlsManagerCacheDir").(string))
		engine.AutoTLSManager.Prompt = autocert.AcceptTOS
		engine.AutoTLSManager.Email = ctx.Value("TlsManagerEmail").(string)
	}

	engine.Pre(middlewares.RemoveTrailingSlash())
	engine.Pre(middlewares.RequestId())
	engine.Use(middlewares.Recover())
	engine.Use(middlewares.Logger())
	engine.Use(prometheus.Middleware("api"))

	engine.GET("/_metrics", prometheus.Handler())
	engine.GET("/", func(c RequestContext) error {
		return c.JSONPretty(
			http.StatusOK,
			appInfo,
			"  ",
		)
	})

	RegisterRoutes(ctx, engine)

	return &App{
		ctx:    ctx,
		di:     di,
		engine: engine,
	}
}

func (app *App) SetListener(listener net.Listener) {
	app.engine.Listener = listener
}

func (app *App) SetServer(server *http.Server) {
	app.engine.Server = server
}

func (app *App) Start() error {
	listenAt := app.ctx.Value("ListenAt").(string)
	sslEnabled := app.ctx.Value("SslEnabled").(bool)

	err := app.startListener(sslEnabled, listenAt)
	if errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (app *App) startListener(sslEnabled bool, listenAt string) error {
	if sslEnabled {
		return app.engine.StartAutoTLS(listenAt)
	}
	return app.engine.Start(listenAt)
}

func (app *App) Stop() error {
	return app.engine.Shutdown(app.ctx)
}
