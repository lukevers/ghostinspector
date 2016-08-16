package ghostinspector

// Viewport is the sizing structure for a Ghost Inspector test to use when
// running. It's also used when mapping JSON data to structs.
type Viewport struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
