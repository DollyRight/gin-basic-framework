package middleware

import (
	"gin-basic-framework/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{global.GLOBAL_VP.GetString("cors.allow-origins")}
	config.AllowMethods = []string{global.GLOBAL_VP.GetString("cors.allow-methods")}
	config.AllowHeaders = []string{global.GLOBAL_VP.GetString("cors.allow-headers")}
	config.ExposeHeaders = []string{global.GLOBAL_VP.GetString("cors.expose-headers")}

	return cors.New(config)
}
