package main

import (
	"flag"
	"log"
)

// Configuration variables
var (
	Method             string
	Url                string
	Timeout            int
	PostFile           string
	ContentType        string
	DisableCompression bool
	DisableKeepalive   bool
	Verbose            bool
)

// ParseFlags parses the command-line flags and validates required options.
func ParseFlags() {
	flag.StringVar(&Method, "method", "GET", "HTTP method, one of GET, POST")
	flag.StringVar(&Url, "url", "", "Target URL (e.g., https://ausopen.com/)")
	flag.IntVar(&Timeout, "timeout", 10, "Max seconds to wait for each response")
	flag.StringVar(&PostFile, "post-file", "", "File containing data to POST (required for POST)")
	flag.StringVar(&ContentType, "content-type", "text/plain", "Content-Type header")
	flag.BoolVar(&DisableCompression, "disable-compression", false, "Disable HTTP compression")
	flag.BoolVar(&DisableKeepalive, "disable-keepalive", false, "Disable HTTP keep-alives")
	flag.BoolVar(&Verbose, "verbose", false, "Enable verbose debug logging")
	flag.Parse()

	if Url == "" {
		log.Fatalln("Error: --url cannot be empty. Please specify a target URL.")
	}
	if Method != "GET" && Method != "POST" {
		log.Fatalln("Error: HTTP method must be either GET or POST.")
	}
}
