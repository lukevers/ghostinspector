package ghostinspector

import (
	"log"
	"testing"
)

func TestSuiteService_List(t *testing.T) {
	response, err := client.Suites.List()
	log.Println(err)
	log.Println(response)
}

func TestSuiteService_Get(t *testing.T) {
	// Using the example suite id from the API docs
	response, err := client.Suites.Get("53322c9fe3db0fcd624cc0a6")
	log.Println(err)
	log.Println(response)
}
