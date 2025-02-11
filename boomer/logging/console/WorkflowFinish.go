package console

import (
	"fmt"
	"time"

	"boomer/model"
)

func WorkflowFinish(meta *model.TestMeta, pass bool) {
	if pass {
		fmt.Printf("%s - %s: Finished [Pass]: TraceId: [%s], SpanId: [%s]\n", time.Now().Format(time.RFC3339Nano), meta.Workflow, meta.TraceId, meta.WorkflowSpanId)
	} else {
		fmt.Printf("%s - %s: Finished [Error]: TraceId: [%s], SpanId: [%s]\n", time.Now().Format(time.RFC3339Nano), meta.Workflow, meta.TraceId, meta.WorkflowSpanId)
	}
}
