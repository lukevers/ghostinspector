package ghostinspector

import (
	"errors"
	"fmt"
)

type Suite struct {
	Gi           *GhostInspector
	Id           string       `json:"_id"`
	Name         string       `json:"name"`
	Organization Organization `json:"organization"`
	TestCount    int          `json:"testCount"`
}

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
			Id:   s["_id"].(string),
			Name: s["name"].(string),
			Organization: Organization{
				Id:   o["_id"].(string),
				Name: o["name"].(string),
			},
			TestCount: int(s["testCount"].(float64)),
		})
	}

	return &suites, nil
}

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
		Id:   s["_id"].(string),
		Name: s["name"].(string),
		Organization: Organization{
			Id:   o["_id"].(string),
			Name: o["name"].(string),
		},
		TestCount: int(s["testCount"].(float64)),
	}

	return &suite, nil
}

func (s *Suite) Execute(startUrl string, immediate bool, viewport *Viewport, userAgent string, vars string) (interface{}, error) {
	args := fmt.Sprintf("startUrl=%s&immediate=%b", startUrl, immediate)
	if viewport != nil {
		args = fmt.Sprintf("%s&viewport=%dx%d", args, viewport.Width, viewport.Width)
	}

	if userAgent != "" {
		args = fmt.Sprintf("%s&userAgent=%s", args, userAgent)
	}

	if vars != "" {
		args = fmt.Sprintf("%s&%s", args, vars)
	}

	response, err := s.Gi.Get(fmt.Sprintf("/suites/%s/execute/", s.Id), args)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (s *Suite) ListTests() (interface{}, error) {
	response, err := s.Gi.Get(fmt.Sprintf("/suites/%s/tests/", s.Id), "")
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
