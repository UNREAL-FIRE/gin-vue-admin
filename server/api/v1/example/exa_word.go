package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
)

type WordApi struct{}

func (e *WordApi) createWord(c *gin.Context) {

	global.GVA_LOG.Debug("添加单词")
}

func (e *WordApi) getOneWord(c *gin.Context) {

}
