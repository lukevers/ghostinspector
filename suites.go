package ghostinspector

import (
	"errors"
	"fmt"
)

// Suite contains information regarding one specific suite.
type Suite struct {
	Gi           *GhostInspector
	ID           string       `json:"_id"`
	Name         string       `json:"name"`
	Organization Organization `json:"organization"`
	TestCount    int          `json:"testCount"`
}

// ListSuites finds all suites available and returns a slice of them.
func (gi *GhostInspector) ListSuites() (*[]Suite, error) {
	response, err := gi.Get("/suites/", "")
	if err != nil {
		return nil, err
	}

	if response.Data == nil {
		return nil, errors.New("Could not get suites")
	}

	var suites []Suite
	for _, suite := range response.Data.([]interface{}) {
		s := suite.(map[string]interface{})
		o := s["organization"].(map[string]interface{})

		suites = append(suites, Suite{
			Gi:   gi,
			ID:   s["_id"].(string),
			Name: s["name"].(string),
			Organization: Organization{
				ID:   o["_id"].(string),
				Name: o["name"].(string),
			},
			TestCount: int(s["testCount"].(float64)),
		})
	}

	return &suites, nil
}

// GetSuite searches for one specific suite and returns the suite if it exists.
// If it does not exists, nil is returned for the suite and an error is also
// returned.
func (gi *GhostInspector) GetSuite(id string) (*Suite, error) {
	response, err := gi.Get(fmt.Sprintf("/suites/%s", id), "")
	if err != nil {
		return nil, err
	}

	if response.Data == nil {
		return nil, errors.New("Could not get suite")
	}

	s := response.Data.(map[string]interface{})
	o := s["organization"].(map[string]interface{})

	suite := Suite{
		Gi:   gi,
		ID:   s["_id"].(string),
		Name: s["name"].(string),
		Organization: Organization{
			ID:   o["_id"].(string),
			Name: o["name"].(string),
		},
		TestCount: int(s["testCount"].(float64)),
	}

	return &suite, nil
}

// Execute runs all tests in the suite.
func (s *Suite) Execute(startURL string, immediate bool, viewport *Viewport, userAgent string, vars string) (interface{}, error) {
	args := fmt.Sprintf("startUrl=%s&immediate=%b", startURL, immediate)
	if viewport != nil {
		args = fmt.Sprintf("%s&viewport=%dx%d", args, viewport.Width, viewport.Width)
	}

	if userAgent != "" {
		args = fmt.Sprintf("%s&userAgent=%s", args, userAgent)
	}

	if vars != "" {
		args = fmt.Sprintf("%s&%s", args, vars)
	}

	response, err := s.Gi.Get(fmt.Sprintf("/suites/%s/execute/", s.ID), args)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// ListTests finds all tests that belong to the suite checked on.
func (s *Suite) ListTests() (interface{}, error) {
	response, err := s.Gi.Get(fmt.Sprintf("/suites/%s/tests/", s.ID), "")
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
