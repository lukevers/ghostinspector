package ghostinspector

// Response is the overlaying structure that is sent back from the Ghost
// Inspector API, and used when mapping JSON data to structs.
type Response struct {
	Code      string      `json:"code"`
	Data      interface{} `json:"data"`
	ErrorType string      `json:"errorType"`
	Message   string      `json:"message"`
}
