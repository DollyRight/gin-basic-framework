package v1

import (
	"github.com/gin-gonic/gin"
)

func InitBaoxinAnomalyDetectRouter(Router *gin.RouterGroup) {
	baoxinAnomalyDetectRouter := Router.Group("baoxin")
	{
		InitRuleRouter(baoxinAnomalyDetectRouter)
	}

}
