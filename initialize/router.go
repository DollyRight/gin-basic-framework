package initialize

import (
	"fmt"
	_ "gin-basic-framework/docs"
	"gin-basic-framework/global"
	"gin-basic-framework/middleware"
	v1 "gin-basic-framework/routers/api/v1"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	router := gin.New()
	////配置日志器
	//router.Use(gin.Logger())

	router.Use(gin.Recovery())
	//swagger文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// jwt middleware
	authMiddleware, err := middleware.JwtMiddleware()
	if err != nil {
		global.GLOBAL_LOGGER.Fatal("JWT Error:" + err.Error())
	}
	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		global.GLOBAL_LOGGER.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		global.GLOBAL_LOGGER.Info(fmt.Sprintf("NoRoute claims: %#v\n", claims))
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// cors跨域
	router.Use(cors.Default())
	//router.Use(middleware.CORS())

	router.POST("/login", authMiddleware.LoginHandler)
	//加载路由组
	apiv1 := router.Group("/api/v1")
	apiv1.GET("/refresh_token", authMiddleware.RefreshHandler)
	//apiv1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.InitApiV1Router(apiv1)
	}

	return router
}
