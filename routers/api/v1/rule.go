package v1

import (
	v1 "gin-basic-framework/api/v1"
	"github.com/gin-gonic/gin"
)

type RuleRouter struct{}

func InitRuleRouter(Router *gin.RouterGroup) {
	RuleRouter := Router.Group("rule")
	ruleApi := v1.ApiGroupApp.RuleApi
	{
		RuleRouter.POST("/create", ruleApi.CreateRuleTask) // 创建用户
		RuleRouter.GET("/stop", ruleApi.StopRuleTask)      // 获取用户
	}
}
