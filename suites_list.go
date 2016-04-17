package ghostinspector

import (
	"log"
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
func (s *SuiteService) List() *SuiteListResponse {
	res := SuiteListResponse{}
	_, err, er := s.client.Get("/suites", []string{}, &res)
	if err != nil {
		log.Println(err)
		log.Println(er)
	}

	log.Println(res)

	return &res
}
