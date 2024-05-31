package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type WordPage struct {
	request.PageInfo        // 分页数据
	Library          string `json:"library"` // 词库
}
