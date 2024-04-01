package main

import (
	"context"
	"fmt"
	"maktapp-api/cmd/api/config"
	"maktapp-api/internal/apps/api"
	"maktapp-api/internal/di"
	"maktapp-api/internal/handlers"
)

var (
	diContainer di.Container

	app *api.App
)

func main() {
	httpConfig := config.Http()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "AppName", "maktapp/api")
	ctx = context.WithValue(ctx, "AppVersion", "1.0.0")
	ctx = context.WithValue(ctx, "TlsManagerEmail", httpConfig.TlsManager.Email)
	ctx = context.WithValue(ctx, "TlsManagerHostWhitelist", httpConfig.TlsManager.HostWhitelist)
	ctx = context.WithValue(ctx, "TlsManagerCacheDir", httpConfig.TlsManager.CacheDir)
	ctx = context.WithValue(ctx, "SslEnabled", httpConfig.Ssl.Enabled)
	ctx = context.WithValue(ctx, "ListenAt", fmt.Sprintf("%v:%s", httpConfig.Listener.Hostname, httpConfig.Listener.Port))

	diContainer = di.NewDI(ctx)
	app = api.NewApp(ctx, diContainer)

	registerDependencies(diContainer)

	handlers.StartLifecycle(app)
}

func registerDependencies(di di.Container) {
	di.MustSet("app", app)
}
