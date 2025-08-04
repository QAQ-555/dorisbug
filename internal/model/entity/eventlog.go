// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Eventlog is the golang structure for table eventlog.
type Eventlog struct {
	EventTime *gtime.Time `json:"eventTime" orm:"event_time" description:""` //
	Adid      string      `json:"adid"      orm:"adid"       description:""` //
	AdId      string      `json:"adId"      orm:"ad_id"      description:""` //
}
