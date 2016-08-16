package ghostinspector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_V1_PATH = "https://api.ghostinspector.com/v1"

type GhostInspector struct {
	apikey string
}

func New(apikey string) *GhostInspector {
	return &GhostInspector{
		apikey: apikey,
	}
}

func (gi *GhostInspector) Get(path, args string) (*Response, error) {
	response, err := http.Get(fmt.Sprintf("%s%s?apiKey=%s&%s", API_V1_PATH, path, gi.apikey, args))
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
