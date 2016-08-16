package ghostinspector

type Response struct {
	Code      string      `json:"code"`
	Data      interface{} `json:"data"`
	ErrorType string      `json:"errorType"`
	Message   string      `json:"message"`
}
