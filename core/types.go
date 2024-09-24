package core

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Response struct {
	Body       string
	StatusCode int
	Headers    http.Header
	Err        error
}

type Config struct {
	BaseUrl      string
	CorsOrigins  []string
	SSLConfig    *tls.Config
	Timeout      time.Duration
	ExtraHeaders map[string]string
}

type HttpClient struct {
	client  *http.Client
	config  Config
	baseUrl string
}

type AdditionalConfig struct {
	Headers map[string]string
}

const (
	DefaultTimeout   = 10 * time.Second
	DefaultUserAgent = "CustomGoHttpClient"
)
