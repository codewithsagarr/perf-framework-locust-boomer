package main

import (
	"fmt"
	"time"

	// "github.com/macnak/boomer/boomer"
	// "github.com/macnak/boomer/otel"
	// "github.com/macnak/boomer/utils"

	// otel2 "go.opentelemetry.io/otel"
	"boomer/boomer"
)

func main() {
	time.Sleep(10 * time.Second)
	fmt.Println("Boomer Starting")
	// providerShutdown := otel.Configure(otel.GetEnvOtel("OTEL_COLLECTOR"), utils.GetEnv("BOOMER_SERVICE", "testing.unset"), utils.GetEnv("ENVIRONMENT", "unset"), utils.GetEnv("RUN_ID", "unset"))
	// if providerShutdown != nil {
	// 	defer providerShutdown()
	// }
	// tracer := otel2.Tracer("Boomer")
	// _, span := tracer.Start(context.Background(), "Boomer")
	// if span != nil {
	// 	defer span.End()
	// }

	boomer.InitBoomer()
}
