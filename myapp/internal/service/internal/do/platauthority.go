// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-09-19 15:24:37
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Platauthority is the golang structure of table platauthority for DAO operations like Where/Data.
type Platauthority struct {
	g.Meta   `orm:"table:platauthority, do:true"`
	Platflag interface{} // 渠道名称_英文
	Platcid  interface{} // 大平台
	Platnid  interface{} // 渠道ID
}
