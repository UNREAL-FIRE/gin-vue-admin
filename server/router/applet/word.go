package applet

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WordRouter struct{}

func (e *WordRouter) InitWordRouter(Router *gin.RouterGroup) {
	wordRouterWithoutRecord := Router.Group("word")
	wordApi := v1.ApiGroupApp.AppletApiGroup.WordApi
	{
		wordRouterWithoutRecord.POST("wordPage", wordApi.WordPage) // 根据词库和章节获取20个单词
	}
}
