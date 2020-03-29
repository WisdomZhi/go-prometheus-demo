package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"go-prometheus-demo/go-micro/pkg/constant"
	demoService "go-prometheus-demo/go-micro/demo-service/proto"
	"net/http"
)

type Demo struct {
	demoService demoService.DemoService
}

func NewDemo(client client.Client) *Demo {
	demoService := demoService.NewDemoService(constant.DemoServiceName, client)
	return &Demo{demoService:demoService}
}

func (d *Demo) SayHello(c *gin.Context)  {
	req := &demoService.GetUsernameReq{}
	res, err := d.demoService.GetUsername(context.TODO(), req)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusOK, "inner err")
		return
	}
	c.JSON(http.StatusOK, "hello,"+res.Username)
}