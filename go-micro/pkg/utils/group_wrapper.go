package utils

import (
	"github.com/gin-gonic/gin"
	"go-prometheus-demo/go-micro/pkg/constant"
)

type GroupWrapper struct {
	group *gin.RouterGroup
}

func NewGroupWrapper(group *gin.RouterGroup) *GroupWrapper {
	group.Handlers = append([]gin.HandlerFunc{func(context *gin.Context) {

	}}, group.Handlers...)
	return &GroupWrapper{
		group: group,
	}
}

func (tg *GroupWrapper) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	tg.wrapperHandler(relativePath)
	return tg.group.GET(relativePath, handlers...)
}

func (tg *GroupWrapper) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	tg.wrapperHandler(relativePath)
	return tg.group.POST(relativePath, handlers...)
}

func (tg *GroupWrapper) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	tg.wrapperHandler(relativePath)
	return tg.group.DELETE(relativePath, handlers...)
}

func (tg *GroupWrapper) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	tg.wrapperHandler(relativePath)
	return tg.group.PUT(relativePath, handlers...)
}

func (tg *GroupWrapper) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	tg.wrapperHandler(relativePath)
	return tg.group.PATCH(relativePath, handlers...)
}

func (tg *GroupWrapper) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	tg.wrapperHandler(relativePath)
	return tg.group.OPTIONS(relativePath, handlers...)
}

func (tg *GroupWrapper) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	tg.wrapperHandler(relativePath)
	return tg.group.HEAD(relativePath, handlers...)
}

func (tg *GroupWrapper) Use(handlers ...gin.HandlerFunc) {
	tg.group.Use(handlers...)
}

func (tg *GroupWrapper) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	tg.wrapperHandler(relativePath)
	return tg.group.Any(relativePath, handlers...)
}

func (tg *GroupWrapper) Handle(httpMethod string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	tg.wrapperHandler(relativePath)
	return tg.group.Handle(httpMethod, relativePath, handlers...)
}

func (tg *GroupWrapper) Group(relativePath string, handlers ...gin.HandlerFunc) *GroupWrapper {
	group := tg.group.Group(relativePath, handlers...)
	return NewGroupWrapper(group)
}

//func (tg *GroupWrapper) wrapperHandler(relativePath string, handlers ...gin.HandlerFunc) []gin.HandlerFunc {
//	handlers = append([]gin.HandlerFunc{func(context *gin.Context) {
//		context.Set(RelativePathKey, tg.group.BasePath()+relativePath)
//	}}, handlers...)
//	return handlers
//}

//保存相对路径到context中
func (tg *GroupWrapper) wrapperHandler(relativePath string) {
	if len(tg.group.Handlers) > 0 {
		tg.group.Handlers[0] = func(context *gin.Context) {
			context.Set(constant.RelativePathKey, tg.group.BasePath()+relativePath)
		}
	}
}
