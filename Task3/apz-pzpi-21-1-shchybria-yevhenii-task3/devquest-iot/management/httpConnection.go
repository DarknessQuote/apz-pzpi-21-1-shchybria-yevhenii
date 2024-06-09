package management

import (
	"net/http"
)

type HttpConnection struct {
	Client *http.Client
	serverHost string
}

func NewHttpConnection(config DeviceConfig) *HttpConnection {
	tr := &http.Transport{
		MaxIdleConns: config.ConnectionSettings.MaxIdleConns,
		IdleConnTimeout: config.ConnectionSettings.IdleConnTimeout,
		DisableCompression: config.ConnectionSettings.DisableCompression,
	}

	return &HttpConnection{
		Client: &http.Client{Transport: tr, Timeout: config.ConnectionSettings.ConnTimeout},
		serverHost: config.ConnectionSettings.ServerHost,
	}
}