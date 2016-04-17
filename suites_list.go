package ghostinspector

import (
	"time"
)

type SuiteListResponse struct {
	Code string `json:"code"`
	Data []struct {
		ID           string `json:"_id"`
		Organization struct {
			ID   string `json:"_id"`
			Name string `json:"name"`
		} `json:"organization"`
		DateCreated time.Time `json:"dateCreated"`
		TestCount   int       `json:"testCount"`
		Name        string    `json:"name"`
	} `json:"data"`
}

// Fetch an array of all the suites in your account.
//
// https://ghostinspector.com/docs/api/suites/#list
func (s *SuiteService) List() (*SuiteListResponse, error) {
	res := SuiteListResponse{}
	_, err := s.client.Get("/suites", []string{}, &res)
	return &res, err
}
