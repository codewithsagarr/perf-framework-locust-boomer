package console

import (
	"fmt"
	"time"

	"boomer/model"
)

func WorkflowStart(meta *model.TestMeta) {
	fmt.Printf("%s - %s: Started\n", time.Now().Format(time.RFC3339Nano), meta.Workflow)
}
