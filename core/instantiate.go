package core

import (
	"crypto/tls"
	"net/http"
)

func NewClient(baseUrl string) *HttpClient {
	defaultConfig := Config{
		BaseUrl:      baseUrl,
		CorsOrigins:  []string{"*"},
		SSLConfig:    &tls.Config{InsecureSkipVerify: true},
		Timeout:      DefaultTimeout,
		ExtraHeaders: map[string]string{"User-Agent": DefaultUserAgent},
	}

	transport := &http.Transport{
		TLSClientConfig: defaultConfig.SSLConfig,
	}

	return &HttpClient{
		client: &http.Client{
			Timeout:   defaultConfig.Timeout,
			Transport: transport,
		},
		config:  defaultConfig,
		baseUrl: baseUrl,
	}
}

func NewClientWithConfig(config Config) *HttpClient {
	transport := &http.Transport{
		TLSClientConfig: config.SSLConfig,
	}

	return &HttpClient{
		client: &http.Client{
			Timeout:   config.Timeout,
			Transport: transport,
		},
		config:  config,
		baseUrl: config.BaseUrl,
	}
}
