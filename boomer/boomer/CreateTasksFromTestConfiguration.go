package boomer

import (
	"boomer/model"

	"boomer/boomer/workflow"

	"github.com/myzhan/boomer"
)

func CreateTasksFromTestConfiguration(tc model.TestConfiguration) []*boomer.Task {
	var tasks []*boomer.Task
	for _, task := range tc.Tasks {
		newTask := &boomer.Task{
			Name:   task.Name,
			Weight: int(task.Weight),
			Fn:     workflow.GetWorkflow(task.Name),
		}
		tasks = append(tasks, newTask)
	}
	return tasks
}
