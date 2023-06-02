package router

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/hole/api/resthandler"
)


type UserRouter interface {
	InitUserRouter(userRouter *gin.RouterGroup)
}

type UserRouterImpl struct {
	UserRestHandler resthandler.UserRestHandler
}

func NewUserRouterImpl(userRestHandler resthandler.UserRestHandler) *UserRouterImpl {
	return &UserRouterImpl{UserRestHandler: userRestHandler}
}

func (impl UserRouterImpl) InitUserRouter(userRouter *gin.RouterGroup) {
	// 获取登录表单页面
	userRouter.GET("/login", impl.UserRestHandler.ShowLogin)

	// 提交登录表单

}
