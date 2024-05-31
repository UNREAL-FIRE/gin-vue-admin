package applet

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applet"
	appletReq "github.com/flipped-aurora/gin-vue-admin/server/model/applet/request"
)

type WordService struct{}

func (wordService *WordService) CreateExaWord(e applet.EnglishWord) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (wordService *WordService) GetOneWord(id uint) (word applet.EnglishWord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&word).Error
	return
}

func (wordService *WordService) WordPage(search appletReq.WordPage) (list []applet.EnglishWord, total int64, err error) {
	limit := search.PageSize
	offset := search.PageSize * (search.Page - 1)
	db := global.GVA_DB.Model(&applet.EnglishWord{})
	var entities []applet.EnglishWord
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Limit(limit).Offset(offset).Where("library = ?", search.Library).Order("id asc").Find(&entities).Error
	return entities, total, err
}
