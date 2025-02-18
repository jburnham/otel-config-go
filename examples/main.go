package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"

	"github.com/honeycombio/otel-config-go/otelconfig"
)

func main() {
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry()

	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}

	defer otelShutdown()
	tracer := otel.Tracer("my-app")
	ctx := context.Background()
	ctx, span := tracer.Start(ctx, "doing-things")
	defer span.End()
}
