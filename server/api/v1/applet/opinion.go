package applet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OpinionApi struct {
}

func (e *OpinionApi) CreateOpinion(c *gin.Context) {
	var opinion applet.Opinion
	err := c.ShouldBindJSON(&opinion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(opinion, utils.OpinionVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = OpinionService.CreateOpinion(opinion)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
