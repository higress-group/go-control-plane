module github.com/envoyproxy/go-control-plane/envoy

go 1.17

// Used to resolve import issues related to go-control-plane package split (https://github.com/envoyproxy/go-control-plane/issues/1074)
replace github.com/envoyproxy/go-control-plane => ../

require (
	github.com/census-instrumentation/opencensus-proto v0.4.1
	github.com/cncf/xds/go v0.0.0-20250326154945-ae57f3c0d45f
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/golang/protobuf v1.5.4
	github.com/prometheus/client_model v0.6.2
	go.opentelemetry.io/proto/otlp v1.7.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250528174236-200df99c418a
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a
	google.golang.org/grpc v1.73.0
	google.golang.org/protobuf v1.36.6
)

require (
	cel.dev/expr v0.23.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
)
