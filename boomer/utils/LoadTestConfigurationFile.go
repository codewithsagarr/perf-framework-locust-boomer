package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"boomer/model"
)

// Used to load the test configuration file that contains all the tasks that are to be run for the test.
func LoadTestConfigurationFile(path string) model.TestConfiguration {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var testConfiguration model.TestConfiguration
	if err := json.Unmarshal(content, &testConfiguration); err != nil {
		log.Fatal(err)
	}

	return testConfiguration
}
