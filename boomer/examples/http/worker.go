package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/myzhan/boomer"
)

// WorkerTask is the function that will be executed as a load-test task.
func WorkerTask() {
	// Create a new HTTP request using the configured method and URL.
	req, err := http.NewRequest(Method, Url, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", ContentType)

	// Record the start time.
	startTime := time.Now()
	resp, err := client.Do(req)
	elapsed := time.Since(startTime).Nanoseconds() / int64(time.Millisecond)

	// If an error occurs, record a failure.
	if err != nil {
		if Verbose {
			log.Printf("Request error: %v", err)
		}
		boomer.RecordFailure("http", "error", elapsed, err.Error())
		return
	}
	defer resp.Body.Close()

	// If verbose logging is enabled, print the response body.
	if Verbose {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response: %v", err)
		} else {
			fmt.Printf("Status Code: %d\n", resp.StatusCode)
			fmt.Printf("Response Body: %s\n", string(body))
		}
	} else {
		// Otherwise, discard the response body.
		io.Copy(io.Discard, resp.Body)
	}

	// Record the success with response time and content length.
	boomer.RecordSuccess("http", strconv.Itoa(resp.StatusCode), elapsed, resp.ContentLength)
}
