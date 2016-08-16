package ghostinspector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIV1Path is the protocol and base path of the Ghost Inspector API version 1
const APIV1Path = "https://api.ghostinspector.com/v1"

// GhostInspector handles sending to the API
type GhostInspector struct {
	apikey string
}

// New creates a new Ghost Inspector client that can be used to send requests.
func New(apikey string) *GhostInspector {
	return &GhostInspector{
		apikey: apikey,
	}
}

// Get sends GET requests to the Ghost Inspector API, using APIV1Path as the
// base path. This is mainly to be used internally, but exported for users if
// new API calls exists without them being reflected in this package.
func (gi *GhostInspector) Get(path, args string) (*Response, error) {
	response, err := http.Get(fmt.Sprintf("%s%s?apiKey=%s&%s", APIV1Path, path, gi.apikey, args))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resp := Response{}
	json.Unmarshal(body, &resp)
	return &resp, nil
}
