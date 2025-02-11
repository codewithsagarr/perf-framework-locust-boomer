package utils

import (
	"crypto/tls"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func GetHTTPClient(https bool, sslverify bool) *http.Client {

	if https {
		if sslverify {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
		} else {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
	}
	tr := otelhttp.NewTransport(
		http.DefaultTransport,
	)
	return &http.Client{
		Timeout:   120 * time.Second,
		Transport: tr,
	}
}
