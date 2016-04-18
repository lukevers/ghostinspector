package ghostinspector

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
