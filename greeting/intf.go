package greeting

import (
	"context"

	"go-pb-swagger/hello"
)

type GreetingService interface {
	Hello(ctx context.Context, req hello.HelloRequest) (*hello.HelloResponse, error)
}
