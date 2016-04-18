package ghostinspector

import (
	"fmt"
)

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
