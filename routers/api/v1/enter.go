package v1

import (
	"github.com/gin-gonic/gin"
)

type ApiV1Router struct{}

func InitApiV1Router(Router *gin.RouterGroup) {
	apiv1Router := Router.Group("")
	{
		InitUserRouter(apiv1Router)
		InitBaoxinAnomalyDetectRouter(apiv1Router)
	}
}
