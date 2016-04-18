package ghostinspector

import (
	"fmt"
	"time"
)

type SuiteService struct {
	client *Client
}

type Suite struct {
	ID           string       `json:"_id"`
	Organization Organization `json:"organization"`
	DateCreated  time.Time    `json:"dateCreated"`
	TestCount    int          `json:"testCount"`
	Name         string       `json:"name"`
}

// The response struct for calling *SuiteService.List
type SuiteListResponse struct {
	Code string  `json:"code"`
	Data []Suite `json:"data"`
}

// Fetch an array of all the suites in your account.
//
// https://ghostinspector.com/docs/api/suites/#list
func (s *SuiteService) List() (*SuiteListResponse, error) {
	res := SuiteListResponse{}
	_, err := s.client.Get("/suites", []string{}, &res)
	return &res, err
}

// The response struct when calling *SuiteService.Get
type SuiteGetResponse struct {
	Code string `json:"code"`
	Data Suite  `json:"data"`
}

// Fetch a single suite from your account.
//
// https://ghostinspector.com/docs/api/suites/#get
func (s *SuiteService) Get(id string) (*SuiteGetResponse, error) {
	res := SuiteGetResponse{}
	_, err := s.client.Get(
		fmt.Sprintf("/suites/%s/", id),
		[]string{},
		&res,
	)

	return &res, err
}
