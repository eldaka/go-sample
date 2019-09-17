package api

import (
	"net/http"
	"time"

	config "github.com/eldaka/go-sample/config/xfers"
)

var (
	clientReq  *http.Client
	mainConfig config.Config
)

//InitHTTPClientRequest represents to initialize http client wrapper
func InitHTTPClientRequest(cfg *config.Config) {
	mainConfig = *cfg

	// TODO better to wrap http.Client with our own struct
	// that way we can add circuit breaker + prometheus in all of our API call to 3rd party API
	clientReq = &http.Client{Timeout: mainConfig.Server.HTTPTimeout * time.Second}
}
