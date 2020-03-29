package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
	"github.com/micro/go-micro/web"
	"go-prometheus-demo/go-micro/demo-api/handler"
	"go-prometheus-demo/go-micro/pkg/constant"
	"go-prometheus-demo/go-micro/pkg/utils"
)

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	service := web.NewService(
		web.Context(ctx),
		web.Name(constant.DemoApiName),
		web.RegisterTTL(constant.RegisterTTL),
		web.RegisterInterval(constant.RegisterInterval),
	)

	service.Init(
		web.Action(func(i *cli.Context) {
			InitHandler(service)
		}),
	)

	if err := service.Run(); err != nil {
		fmt.Println("service run error")
	}
	cancel()
}

func InitHandler(service web.Service)  {
	c := service.Options().Service.Client()
	demo := handler.NewDemo(c)
	router := gin.Default()

	//添加prometheus中间件
	promMonitor := utils.NewPrometheusMonitor(constant.DemoApiNameSpace, constant.DemoApiName)
	router.Use(promMonitor.PromMiddleware())

	g := router.Group("/demo-api")
	//有rest风格路径时可以使用这个封装的GroupWrapper代替使用
	//g := utils.NewGroupWrapper(router.Group("/demo-api"))

	g.GET("/sayHello", demo.SayHello)
	service.Handle("/", router)

	go utils.StartMonitor("127.0.0.1", 8887)
}