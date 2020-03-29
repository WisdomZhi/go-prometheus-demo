package main

import (
	"context"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"go-prometheus-demo/go-micro/demo-service/handler"
	demoService "go-prometheus-demo/go-micro/demo-service/proto"
	"go-prometheus-demo/go-micro/pkg/constant"
	"go-prometheus-demo/go-micro/pkg/utils"
	"os"
)

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	service := micro.NewService(
		micro.RegisterTTL(constant.RegisterTTL),
		micro.RegisterInterval(constant.RegisterInterval),
		micro.Context(ctx),
		micro.Name(constant.DemoServiceName),

		//添加prometheus插件
		micro.WrapHandler(prometheus.NewHandlerWrapper(
			server.Name(constant.DemoServiceName)),
		),
	)

	service.Init(micro.Action(func(i *cli.Context) {
		err := InitHandler(service)
		if  err != nil {
			fmt.Println("InitHandler error:", err)
			os.Exit(0)
		}
	}))

	if err := service.Run(); err != nil {
		fmt.Println("service run error:", err)
	}
	cancel()
}

func InitHandler(service micro.Service) error {
	demo := handler.NewDemo()
	err := demoService.RegisterDemoServiceHandler(service.Server(), demo)
	if err != nil {
		return err
	}

	go utils.StartMonitor("127.0.0.1", 8888)
	return err
}
