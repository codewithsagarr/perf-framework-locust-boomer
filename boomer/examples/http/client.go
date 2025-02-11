package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

// NewHTTPClient creates and returns a configured HTTP client.
func NewHTTPClient(timeoutSeconds int, disableCompression, disableKeepalive bool) *http.Client {
	// Create a custom transport with desired settings.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // Warning: Disables certificate verification
		},
		MaxIdleConnsPerHost: 2000,
		DisableCompression:  disableCompression,
		DisableKeepAlives:   disableKeepalive,
	}

	return &http.Client{
		Transport: tr,
		Timeout:   time.Duration(timeoutSeconds) * time.Second,
	}
}
