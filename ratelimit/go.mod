module github.com/envoyproxy/go-control-plane/ratelimit

go 1.17

replace github.com/envoyproxy/go-control-plane/envoy => ../envoy

require (
	github.com/envoyproxy/go-control-plane/envoy v1.32.4
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/cncf/xds/go v0.0.0-20230607035331-e9ce68804cb4 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
