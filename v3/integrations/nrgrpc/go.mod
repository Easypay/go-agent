module github.com/Easypay/go-agent/v3/integrations/nrgrpc

// As of Dec 2019, the grpc go.mod file uses 1.11:
// https://github.com/grpc/grpc-go/blob/master/go.mod
go 1.11

require (
	github.com/Easypay/go-agent/v3 v3.3.9
	// protobuf v1.3.0 is the earliest version using modules, we use v1.3.1
	// because all dependencies were removed in this version.
	github.com/golang/protobuf v1.3.3
	// v1.15.0 is the earliest version of grpc using modules.
	google.golang.org/grpc v1.27.0
)
