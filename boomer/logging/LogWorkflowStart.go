package logging

import (
	"boomer/logging/console"
	"boomer/model"
)

func LogWorkflowStart(meta *model.TestMeta) {
	console.WorkflowStart(meta)
}
