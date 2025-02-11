package workflow

import (
	"time"

	"boomer/boomer/endpoints/http_server"
	"boomer/logging"
	"boomer/logging/console"
	"boomer/model"

	// "github.com/macnak/boomer/endpoints/grpc_server"
	// "github.com/macnak/boomer/endpoints/http_server"

	"boomer/utils"
)

func Workflow_02() {
	var hasError = false

	meta := &model.TestMeta{
		Workflow: "02_Workflow",
		Step:     0,
		Label:    "",
		RunId:    utils.GetEnv("RUN_ID", "unset"),
		User:     "+61411000000",
	}
	console.WorkflowStart(meta)

	// otel.NewTrace(meta)

	// otel.NewWorkflowSpan(meta)
	boomerStop := EventSubscription(meta)

	WorkflowNextStep(meta, "GetAusOpen")
	// step_01 := http_server.AddUser(meta)
	step_01 := http_server.GetAusOpen(meta)
	logging.LogWorkflowStep(meta, step_01)
	hasError = !step_01.Success
	time.Sleep(10 * time.Second)

	// if !hasError {
	// 	WorkflowNextStep(meta, "gRPCConnection")
	// 	passed := utils.NewTestGRPC(meta, utils.GetEnv("GRPC_END_POINT", "node_grpc_server:50051"), "", false)
	// 	if passed {
	// 		console.WorkflowStep(meta, true, "")
	// 	} else {
	// 		hasError = true
	// 		console.WorkflowStep(meta, false, "")
	// 	}
	// }
	// time.Sleep(1 * time.Second)

	// if !hasError {
	// 	WorkflowNextStep(meta, "New.GetAllNews")
	// 	passed, output := grpc_server.GetAllNews(meta)
	// 	if passed {
	// 		console.WorkflowStep(meta, true, output)
	// 	} else {
	// 		hasError = true
	// 		console.WorkflowStep(meta, false, "")
	// 	}
	// }
	time.Sleep(1 * time.Second)

	console.WorkflowFinish(meta, !hasError)
	EventUnsubscription(meta, boomerStop)
}
