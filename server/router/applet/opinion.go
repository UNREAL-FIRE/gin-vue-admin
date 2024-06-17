package applet

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type OpinionRouter struct {
}

func (e *OpinionRouter) InitOpinionRouter(Router *gin.RouterGroup) {
	OpinionRouterWithoutRecord := Router.Group("opinion")
	opinionApi := v1.ApiGroupApp.AppletApiGroup.OpinionApi
	{
		OpinionRouterWithoutRecord.POST("opinion", opinionApi.createOpinion)
	}
}
