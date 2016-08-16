package ghostinspector

import (
	"fmt"
)

// Test contains information regarding one specific test.
type Test struct {
	Gi   *GhostInspector
	ID   string `json:"_id"`
	Name string `json:"name"`
}

// Execute runs one specific test.
func (t *Test) Execute(startURL string, immediate bool, viewport *Viewport, userAgent string, vars string) (interface{}, error) {
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

	response, err := t.Gi.Get(fmt.Sprintf("/tests/%s/execute/", t.ID), args)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
