package model

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
)

type TestMeta struct {
	RunId            string
	Workflow         string
	WorkflowCtx      context.Context
	WorkflowSpan     trace.Span
	WorkflowSpanId   string
	Step             int
	StepSpanId       string
	Label            string
	User             string
	Summary          bool
	StartTime        time.Time
	EndTime          time.Time
	TotalElapsedTime time.Duration
	TraceId          string
	Tracer           trace.Tracer
	// TestGRPC         *TestGRPC
}
