package greeter

import (
	"context"
	"google.golang.org/grpc"
	gpb "bff/internal/gen/greeter"
)

type GreeterClient interface {
	SayHello(ctx context.Context, in *gpb.HelloRequest, opts ...grpc.CallOption) (*gpb.HelloReply, error)
}

