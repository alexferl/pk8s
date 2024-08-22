// Code generated; DO NOT EDIT.

package k8s

// GRPCActionV1
type GRPCActionV1 struct {
	// Port number of the gRPC service. Number must be in the range 1 to 65535.
	Port int32 `json:"port"`
	// Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.
	Service *string `json:"service,omitempty"`
}
