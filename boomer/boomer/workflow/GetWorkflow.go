package workflow

func GetWorkflow(name string) func() {
	workflows := map[string]func(){
		// "Workflow_01": Workflow_01,
		"Workflow_02": Workflow_02,
		"Workflow_03": Workflow_03,
	}
	return workflows[name]
}
