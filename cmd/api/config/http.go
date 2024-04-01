package config

import (
	"maktapp-api/internal/config"
	"maktapp-api/internal/env"
)

func Http() config.HttpConfig {
	return config.HttpConfig{
		Listener: config.Listener{
			Hostname: env.Get("API_LISTENER_HOSTNAME"),
			Port:     env.Get("API_LISTENER_PORT"),
		},
		Ssl: config.SslSettings{
			Enabled: env.AsBool("API_SSL_ENABLED"),
		},
		TlsManager: config.TlsManager{
			Email:         env.Get("API_TLS_MANAGER_EMAIL"),
			HostWhitelist: env.Get("API_TLS_MANAGER_DOMAIN_WHITELIST"),
			CacheDir:      env.Get("API_TLS_MANAGER_CACHE_DIR"),
		},
	}
}
