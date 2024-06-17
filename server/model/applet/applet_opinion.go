package applet

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Opinion struct {
	global.GVA_MODEL
	Opinion string `json:"opinion" form:"opinion" gorm:"column:opinion;comment:反馈内容"`
}

func (Opinion) TableName() string {
	return "applet_opinion"
}
