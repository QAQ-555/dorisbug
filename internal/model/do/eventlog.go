// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Eventlog is the golang structure of table eventlog for DAO operations like Where/Data.
type Eventlog struct {
	g.Meta    `orm:"table:eventlog, do:true"`
	EventTime *gtime.Time //
	Adid      interface{} //
	AdId      interface{} //
}
