package applet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
	appletReq "github.com/flipped-aurora/gin-vue-admin/server/model/applet/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WordApi struct{}

func (e *WordApi) CreateWord(c *gin.Context) {
	var word applet.EnglishWord
	err := c.ShouldBindJSON(&word)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(word, utils.WordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wordService.CreateExaWord(word)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (e *WordApi) GetOneWord(c *gin.Context) {
	//var word applet.EnglishWord
	//err := c.ShouldBindQuery(&word)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//err = utils.Verify(word.GVA_MODEL, utils.IdVerify)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//data, err := wordService.GetOneWord(word.ID)
	//if err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//	return
	//}
	//response.OkWithDetailed(appletRes.EnglishResponse{EnglishWord: data}, "获取成功", c)
}

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
