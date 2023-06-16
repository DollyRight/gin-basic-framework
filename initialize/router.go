package initialize

import (
	_ "gin-basic-framework/docs"
	"gin-basic-framework/middleware"
	v1 "gin-basic-framework/routers/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	router := gin.New()
	//配置日志器
	router.Use(gin.Logger())

	router.Use(gin.Recovery())
	//swagger文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// cors跨域
	router.Use(cors.Default())
	//router.Use(middleware.CORS())

	//加载路由组
	apiv1 := router.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		v1.InitUserRouter(apiv1)
	}

	return router
}
