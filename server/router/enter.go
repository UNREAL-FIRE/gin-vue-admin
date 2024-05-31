package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/applet"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Applet  applet.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
