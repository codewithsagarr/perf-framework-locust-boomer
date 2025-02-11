package logging

import (
	"fmt"
	"time"

	"boomer/logging/boomer"
	"boomer/logging/console"
	"boomer/model"
)

func LogWorkflowFinish(meta *model.TestMeta, pass bool, elapsed time.Duration, err error) {
	console.WorkflowFinish(meta, pass)

	if pass {
		boomer.BoomerSuccess(meta.Workflow, fmt.Sprintf("%03d_%s", meta.Step, meta.Label), elapsed.Milliseconds(), int64(0))
	} else {
		boomer.BoomerError(meta.Workflow, fmt.Sprintf("%03d_%s", meta.Step, meta.Label), elapsed.Milliseconds(), err.Error())
	}
}
