package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
)

type WordService struct{}

func (exa *WordService) CreateExaWord(e example.ExaWord) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *WordService) GetOneWord(id uint) (word example.ExaWord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&word).Error
	return
}
