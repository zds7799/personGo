// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-09-16 10:18:26
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PrimaryNewrolelog is the golang structure for table primary_newrolelog.
type PrimaryNewrolelog struct {
	Roleid   int64       `json:"roleid"   ` // 角色id
	Serverid int         `json:"serverid" ` // 服务ID
	Platn    string      `json:"platn"    ` // 角色渠道
	Ctime    *gtime.Time `json:"ctime"    ` // 创建时间
	Day      int         `json:"day"      ` // 创建时间day
	Hour     int         `json:"hour"     ` // 创建时间hour
	Month    int         `json:"month"    ` // 创建时间month
}
