package workflow

import (
	"boomer/model"

	"github.com/myzhan/boomer"
)

func EventSubscription(meta *model.TestMeta) func() {
	boomerStop := func() {
		if meta.WorkflowSpan != nil {
			meta.WorkflowSpan.End()
		}
	}
	// boomer.Events.Subscribe

	boomer.Events.Subscribe("boomer:stop", boomerStop)

	return boomerStop
}
