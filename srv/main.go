package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	hello "git.coding.net/wang_zhipeng/micro_demo/proto"
	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/etcd"
)

type Say struct{}

func (s *Say) Call(ctx context.Context, req *hello.CallRequest, rsp *hello.CallResponse) error {
	log.Print("Received Say.Call request")
	rsp.Message = "Hello" + req.GetName()
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.say"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	// hello.RegisterSayHandler(service.Server(), new(Say))
	hello.RegisterExampleHandler(service.Server(), new(Say))

	calls := make([]uintptr, 10)

	n := runtime.Callers(0, calls)
	fmt.Println(n, len(calls))

	for _, c := range calls {
		fmt.Println(runtime.FuncForPC(c).Name())
	}

	fs := runtime.CallersFrames(calls)

	i := 0
	for {
		if f, ok := fs.Next(); ok {
			fmt.Println(f.Function)
		}
		fmt.Println(i)
		i++
		break
	}
	return
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
