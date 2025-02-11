package main

import "net/http"

// Global HTTP client and POST body (if any)
var (
	client   *http.Client
	postBody []byte
)
