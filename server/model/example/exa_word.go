package example

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type ExaWord struct {
	global.GVA_MODEL
	WordName string `json:"wordName" form:"wordName" gorm:"comment:单词"`
	Trans    string `json:"trans" form:"trans" gorm:"comment:翻译"`
	Usphone  string `json:"usphone" form:"usphone" gorm:"comment:美音"`
	Ukphone  string `json:"ukphone" form:"ukphone" gorm:"comment:英音"`
}
