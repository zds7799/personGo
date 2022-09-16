// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-09-16 10:18:26
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PrimaryServerconfig is the golang structure of table serverconfig for DAO operations like Where/Data.
type PrimaryServerconfig struct {
	g.Meta   `orm:"table:serverconfig, do:true"`
	Serverid interface{} // 服务ID
	NameCN   interface{} // 显示名称_中文
	NameEN   interface{} // 显示名称_英文
	Descript interface{} // 描述信息
	Ctime    *gtime.Time // 服务器创建时间
	Valid    interface{} // 是否有效
}
