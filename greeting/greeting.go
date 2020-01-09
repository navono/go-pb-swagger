package greeting

import (
	"context"
	"fmt"

	"go-pb-swagger/hello"
)

type greeting struct{}

func NewGreetingService() GreetingService {
	return &greeting{}
}

// Hello
// @Summary this is summary
// @Description this is description
// @Tags
// @Accept  octet-stream
// @Produce octet-stream
// @Param   request body hello.HelloRequest true "HelloRequest"
// @Success 200 {object} hello.HelloResponse
// @Router / [get]
func (gs *greeting) Hello(ctx context.Context, req hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{
		Msg: fmt.Sprintf("Hello, %s", req.Msg),
	}, nil
}
