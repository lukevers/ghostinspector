package ghostinspector

type Error struct {
	Code      string `json:"code"`
	ErrorType string `json:"errorType"`
	Message   string `json:"message"`
}
