package applet

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Library struct {
	global.GVA_MODEL
	WordName string `json:"wordName" form:"wordName" gorm:"column:word_name;comment:单词"`
	Trans    string `json:"trans" form:"trans" gorm:"column:trans;comment:翻译"`
	Usphone  string `json:"usphone" form:"usphone" gorm:"column:usphone;comment:美音"`
	Ukphone  string `json:"ukphone" form:"ukphone" gorm:"column:ukphone;comment:英音"`
	Library  string `json:"library" form:"library" gorm:"column:library;comment:词库"`
	Notation string `json:"notation" form:"notation" gorm:"column:notation;comment:日语"`
}

func (Library) TableName() string {
	return "applet_library"
}
