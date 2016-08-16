package ghostinspector

import (
	"fmt"
)

type Test struct {
	Gi   *GhostInspector
	Id   string `json:"_id"`
	Name string `json:"name"`
}

func (t *Test) Execute(startUrl string, immediate bool, viewport *Viewport, userAgent string, vars string) (interface{}, error) {
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

	response, err := t.Gi.Get(fmt.Sprintf("/tests/%s/execute/", t.Id), args)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
