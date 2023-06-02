
package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RESTRouter struct {
	logger               *zap.SugaredLogger
	Router               *gin.RouterGroup
	UserRouter           UserRouter
	
	// ...

}

func NewRESTRouter(logger *zap.SugaredLogger, userRouter UserRouter) *RESTRouter {
	return &RESTRouter{
		logger:     logger,
		UserRouter: userRouter,
	}
}



func (r RESTRouter) InitRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")

	// 路由组
	userRouter := v1.Group("/user")
	
	// 中间件
	userRouter.Use()
	
	// 注册路由 
	r.UserRouter.InitUserRouter(userRouter)
}