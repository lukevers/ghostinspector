package ghostinspector

import (
	"testing"
)

func TestSuiteService_List(t *testing.T) {
	client := New(&Config{
		ApiKey: API_KEY,
	})

	client.Suites.List()
}
