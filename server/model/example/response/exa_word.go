package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/example"

type ExaWordResponse struct {
	Word example.ExaWord `json:"word"`
}
