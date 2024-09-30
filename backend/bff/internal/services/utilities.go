package services

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGRPCError(err error) error {
	if grpcStatus, ok := status.FromError(err); ok {
		switch grpcStatus.Code() {
		case codes.InvalidArgument:
			return fmt.Errorf("invalid argument: %v", grpcStatus.Message())
		case codes.NotFound:
			return fmt.Errorf("resource not found: %v", grpcStatus.Message())
		case codes.AlreadyExists:
			return fmt.Errorf("resource already exists: %v", grpcStatus.Message())
		case codes.Unauthenticated:
			return fmt.Errorf("authentication failed: %v", grpcStatus.Message())
		case codes.PermissionDenied:
			return fmt.Errorf("permission denied: %v", grpcStatus.Message())
		case codes.Internal:
			return fmt.Errorf("internal server error: %v", grpcStatus.Message())
		case codes.Unavailable:
			return fmt.Errorf("service unavailable: %v", grpcStatus.Message())
		case codes.DeadlineExceeded:
			return fmt.Errorf("request timeout: %v", grpcStatus.Message())
		default:
			return fmt.Errorf("gRPC error (%s): %v", grpcStatus.Code(), grpcStatus.Message())
		}
	}

	// Return non-gRPC errors as they are
	return err
}
