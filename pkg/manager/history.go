package manager

// history saves the model's version metadata.
// This model exported root directory user selected on each storage service.
type history struct {
	Versions         map[string]int64 `json:"versions"`
	LatestGeneration int64            `json:"latest_generation"`
}
