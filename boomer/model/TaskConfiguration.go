package model

// Used for loading the workflow json files.
type TaskConfiguration struct {
	Name   string `json:"name"`
	Weight int64  `json:"weight"`
}
