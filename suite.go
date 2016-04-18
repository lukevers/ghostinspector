package ghostinspector

import (
	"time"
)

type Suite struct {
	ID           string `json:"_id"`
	Organization struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
	} `json:"organization"`
	DateCreated time.Time `json:"dateCreated"`
	TestCount   int       `json:"testCount"`
	Name        string    `json:"name"`
}
