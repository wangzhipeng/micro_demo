package main

import (
	"context"
	"fmt"

	hello "git.coding.net/wang_zhipeng/micro_demo/proto"
	grpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro/metadata"
	_ "github.com/micro/go-plugins/registry/etcd"
)

func main() {
	service := grpc.NewService()
	service.Init()

	// use the generated client stub
	cl := hello.NewExampleService("go.micro.srv.say", service.Client())

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp, err := cl.Call(ctx, &hello.CallRequest{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Message)
}
