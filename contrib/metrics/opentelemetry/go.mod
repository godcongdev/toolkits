module github.com/origadmin/toolkits/registry/metrics/opencensus

go 1.22.5

replace github.com/origadmin/toolkits => ../../../

require (
	github.com/origadmin/toolkits v0.0.0-00010101000000-000000000000
	go.opencensus.io v0.24.0
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/metric v1.30.0
	go.opentelemetry.io/otel/sdk/metric v1.30.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.opentelemetry.io/otel/sdk v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
)