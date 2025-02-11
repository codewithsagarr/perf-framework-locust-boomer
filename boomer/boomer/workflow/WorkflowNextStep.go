package workflow

import "boomer/model"

func WorkflowNextStep(meta *model.TestMeta, label string) {
	meta.Step = meta.Step + 1
	meta.Label = label
}
