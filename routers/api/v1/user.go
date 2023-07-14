package v1

import (
	v1 "gin-basic-framework/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.UserApi
	{
		userRouter.POST("/", userApi.CreateUser) // 创建用户
		userRouter.GET("/:id", userApi.GetUser)  // 获取用户
	}
}
