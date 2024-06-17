package applet

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	WordApi
	OpinionApi
}

var (
	wordService    = service.ServiceGroupApp.AppletServiceGroup.WordService
	OpinionService = service.ServiceGroupApp.AppletServiceGroup.OpinionService
)
