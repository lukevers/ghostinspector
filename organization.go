package ghostinspector

// Organization contains basic information about an organization, and is used
// when mapping JSON data to structs.
type Organization struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}
