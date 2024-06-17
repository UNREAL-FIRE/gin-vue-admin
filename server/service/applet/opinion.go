package applet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
)

type OpinionService struct {
}

func (op *OpinionService) CreateOpinion(e applet.Opinion) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}
