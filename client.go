package ghostinspector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_BASE_PATH = "https://api.ghostinspector.com/v1"
)

type Client struct {
	client *http.Client
	config *Config

	Suites *SuiteService
}

// A generic response type to check for errors
type Response struct {
	Code string `json:"code"`
}

// Create a new client
func New(config *Config) *Client {
	c := &Client{
		client: http.DefaultClient,
		config: config,
	}

	// Services
	c.Suites = &SuiteService{client: c}

	return c
}

// Send a GET request to Ghost Inspector. The path given is the entire path to
// the response after API_BASE_PATH. The options are optional, of course. The
// response parameter should be an empty struct in the format of the response
// expected. If an error occurs, the error return value will contain the error
// given from the Ghost Inspector API.
func (c *Client) Get(path string, options []string, response interface{}) (*http.Response, error) {
	// Gather options
	opts := "apiKey=" + c.config.ApiKey
	for _, o := range options {
		opts += "&"
		opts += o
	}

	// Send request
	resp, err := c.client.Get(
		fmt.Sprintf(
			"%s%s?%s",
			API_BASE_PATH,
			path,
			opts,
		),
	)

	// Check response
	if err != nil {
		return resp, err
	}

	// Convert
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	// Check response type for ERROR
	r := &Response{}
	var er *Error
	json.Unmarshal(body, &r)
	if r.Code == "ERROR" {
		// If there's an error, unmarshal into an Error struct
		json.Unmarshal(body, &er)
		return resp, errors.New(fmt.Sprintf("%s: %s", er.ErrorType, er.Message))
	} else {
		// If there's no error, unmarshal into the type that was given
		json.Unmarshal(body, &response)
		return resp, nil
	}
}
