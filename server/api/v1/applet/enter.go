package applet

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	WordApi
}

var (
	wordService = service.ServiceGroupApp.AppletServiceGroup.WordService
)
