package model

// Used for loading the workflow json files.
type TestConfiguration struct {
	Description   string              `json:"description"`
	Mode          string              `json:"mode"`
	NumberOfUsers int64               `json:"numberOfUsers,omitempty"`
	SpawnRate     float64             `json:"spawnRate,omitempty"`
	TestDuration  int64               `json:"testDuration,omitempty"`
	CacheData     int64               `json:"cacheData,omitempty"`
	Tasks         []TaskConfiguration `json:"tasks"`
}
