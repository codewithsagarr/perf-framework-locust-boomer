package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/myzhan/boomer"
)

func main() {
	// Parse command-line flags.
	ParseFlags()

	// If method is POST, read the POST payload from the specified file.
	if Method == "POST" {
		if PostFile == "" {
			log.Fatalln("Error: --post-file is required for POST requests.")
		}
		data, err := ioutil.ReadFile(PostFile)
		if err != nil {
			log.Fatalf("Error reading post file: %v", err)
		}
		postBody = data
	}

	// Initialize the HTTP client.
	client = NewHTTPClient(Timeout, DisableCompression, DisableKeepalive)

	// Log the configuration.
	fmt.Printf("Starting load test: %s %s\n", Method, Url)

	// Define a Boomer task that executes WorkerTask.
	task := &boomer.Task{
		Name:   "worker",
		Weight: 10, // Adjust weight as needed.
		Fn:     WorkerTask,
	}

	// Run the Boomer load test.
	// If running in distributed mode, this worker will connect to a Locust master.
	boomer.Run(task)
}
