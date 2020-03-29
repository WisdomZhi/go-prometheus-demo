package handler

import (
	"context"
	service "go-prometheus-demo/go-micro/demo-service/proto"
)

type Demo struct {

}

func NewDemo() *Demo {
	return &Demo{}
}

func (d *Demo) GetUsername(ctx context.Context, in *service.GetUsernameReq, out *service.GetUsernameRes) error {
	out.Username = "jack"
	return nil
}