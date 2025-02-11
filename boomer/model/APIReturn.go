package model

import "time"

type APIReturn struct {
	Success bool
	Elapsed time.Duration
	Error   error
	Input   string
	Output  string
}
