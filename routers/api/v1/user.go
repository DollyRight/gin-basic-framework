package v1

import (
	"gin-basic-framework/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("/", service.CreateUser) // 创建用户
		userRouter.GET("/:id", service.GetUser)  // 获取用户
	}
}
