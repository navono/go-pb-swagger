package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/imroc/req"

	"go-pb-swagger/hello"
)

var (
	url    = "http://localhost:10086/"
	header = req.Header{
		"Content-Type":     "application/octet-stream",
		"X-Requested-With": "XMLHttpRequest",
	}
)

func main() {
	helloReq := &hello.HelloRequest{
		Msg: "John",
	}

	data, err := proto.Marshal(helloReq)
	if err != nil {
		return
	}

	r := req.New()
	resp, err := r.Get(url, header, data)
	if err != nil {
		return
	}

	var helloResp hello.HelloResponse
	err = proto.Unmarshal(resp.Bytes(), &helloResp)
	if err != nil {
		return
	}

	fmt.Println(helloResp.Msg)
}
