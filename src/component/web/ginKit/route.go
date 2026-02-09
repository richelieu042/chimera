package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu042/chimera/v3/src/core/sliceKit"
)

// BindHandlersToRoute
/*
适用场景: 1个路由，n个Method.

@param methods 	nil => 接收所有类型method的请求.	e.g. http.MethodGet、http.MethodPost
@param handlers 其中的元素不能为nil!!!
*/
func BindHandlersToRoute(group gin.IRoutes, route string, methods []string, handlers ...gin.HandlerFunc) {
	handlers = sliceKit.RemoveZeroValues(handlers)
	if len(handlers) == 0 {
		// do nothing
		return
	}

	if len(methods) == 0 {
		// (1) any method
		group.Any(route, handlers...)
		return
	}
	// (2) 特定 methods
	for _, method := range methods {
		group.Handle(method, route, handlers...)
	}
}

// BindHandlersToRoutes 将相同的多个处理器，注册到多个路由.
func BindHandlersToRoutes(group gin.IRoutes, routes []string, methods []string, handlers ...gin.HandlerFunc) {
	routes = sliceKit.Uniq(routes)

	for _, route := range routes {
		BindHandlersToRoute(group, route, methods, handlers...)
	}
}
