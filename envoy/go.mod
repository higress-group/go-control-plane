module github.com/envoyproxy/go-control-plane/envoy

go 1.17

// Used to resolve import issues related to go-control-plane package split (https://github.com/envoyproxy/go-control-plane/issues/1074)
replace github.com/envoyproxy/go-control-plane => ../

require (
	github.com/cncf/xds/go v0.0.0-20250326154945-ae57f3c0d45f
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10
	github.com/prometheus/client_model v0.6.2
	go.opentelemetry.io/proto/otlp v1.7.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250528174236-200df99c418a
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a
	google.golang.org/grpc v1.73.0
	google.golang.org/protobuf v1.36.10
)

require (
	cel.dev/expr v0.25.1 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
)
