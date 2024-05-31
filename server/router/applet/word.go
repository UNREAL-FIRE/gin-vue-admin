package applet

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WordRouter struct{}

func (e *WordRouter) InitWordRouter(Router *gin.RouterGroup) {
	wordRouter := Router.Group("word").Use(middleware.OperationRecord())
	wordRouterWithoutRecord := Router.Group("word")
	wordApi := v1.ApiGroupApp.AppletApiGroup.WordApi
	{
		wordRouter.POST("word", wordApi.CreateWord) // 新建一个单词
	}
	{
		wordRouterWithoutRecord.GET("word", wordApi.GetOneWord)    // 获取单个单词
		wordRouterWithoutRecord.POST("wordPage", wordApi.WordPage) // 根据词库和章节获取20个单词
	}
}
