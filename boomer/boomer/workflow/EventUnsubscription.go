package workflow

import (
	"boomer/model"

	"github.com/myzhan/boomer"
)

func EventUnsubscription(meta *model.TestMeta, event func()) {
	boomer.Events.Unsubscribe("boomer:stop", event)
	// meta.WorkflowSpan.End()
}
