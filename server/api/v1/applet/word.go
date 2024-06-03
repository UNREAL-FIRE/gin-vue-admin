package applet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	appletReq "github.com/flipped-aurora/gin-vue-admin/server/model/applet/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WordApi struct{}

func (e *WordApi) WordPage(c *gin.Context) {
	var search appletReq.WordPage
	err := c.ShouldBindJSON(&search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := wordService.WordPage(search)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     search.Page,
		PageSize: search.PageSize,
	}, "获取成功", c)
}
