package ghostinspector

import (
	"os"
)

var (
	client = New(&Config{
		ApiKey: os.Getenv("API_KEY"),
	})
)
