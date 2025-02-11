package boomer

import (
	"boomer/utils"
	"fmt"
	"os"
)

func InitBoomer() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	} else {
		fmt.Println("Current working directory:", cwd)
	}
	// tc := utils.LoadTestConfigurationFile(fmt.Sprintf("%s/workflowJson/%s", utils.GetEnv("MOUNT_DIR", "/mnt"), utils.GetEnv("TEST_COFIG_FILE", "workflow_1.json")))
	// tc := utils.LoadTestConfigurationFile("workflowJson/workflow_2.json")
	// tc := utils.LoadTestConfigurationFile("boomer/workflowJson/workflow_2.json")
	tc := utils.LoadTestConfigurationFile("/mnt/boomer/workflowJson/workflow_2.json")
	fmt.Println(tc.Description)

	tasks := CreateTasksFromTestConfiguration(tc)
	setBoomerMode(tc, tasks)
}
