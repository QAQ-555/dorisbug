package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestInsertReq struct {
	g.Meta `path:"/test" tags:"test" method:"get" summary:"You first hello api"`
	Adid   string `dc:"adid" json:"adid"`
	AdID   string `dc:"ad_id" json:"ad_id"`
}
type TestInsertRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
