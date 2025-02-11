package logging

import (
	"fmt"
	"time"

	"boomer/model"
)

func LogWorkflowStep(meta *model.TestMeta, apiReturn *model.APIReturn) {
	if apiReturn.Success {
		fmt.Printf("%s - %s: %03d-%s [Pass]: TraceId: [%s], SpanId: [%s]\n", time.Now().Format(time.RFC3339Nano), meta.Workflow, meta.Step, meta.Label, meta.TraceId, meta.StepSpanId)
	} else {
		fmt.Printf("%s - %s: %03d-%s [Error]: TraceId: [%s], SpanId: [%s]\n", time.Now().Format(time.RFC3339Nano), meta.Workflow, meta.Step, meta.Label, meta.TraceId, meta.StepSpanId)
	}
}
