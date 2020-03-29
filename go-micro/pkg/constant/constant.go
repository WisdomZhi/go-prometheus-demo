package constant

import "time"

const (
	DemoApiName = "go-micro.api.demo-api"
	DemoServiceName = "go-micro.demo-service"
	RegisterTTL      = time.Second * 20
	RegisterInterval = time.Second * 10

	RelativePathKey = "relativePathKey"

	DemoApiNameSpace = "demo_api"
)
