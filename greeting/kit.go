package greeting

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"

	"go-pb-swagger/hello"
)

func MakeGreetingEp(s GreetingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(hello.HelloRequest)
		resp, err := s.Hello(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func DecHelloRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil || len(data) == 0 {
		return nil, err
	}

	var helloReq hello.HelloRequest
	if err := proto.Unmarshal(data, &helloReq); err != nil {
		return nil, err
	}

	return helloReq, nil
}

func EncHelloRequest(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(*hello.HelloResponse)
	data, err := proto.Marshal(r)
	if err != nil {
		return err
	}

	header := w.Header()
	if header.Get(echo.HeaderContentType) == "" {
		header.Set(echo.HeaderContentType, echo.MIMEOctetStream)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		return err
	}

	return nil
}
