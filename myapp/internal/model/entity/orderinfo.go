// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-09-19 15:24:37
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Orderinfo is the golang structure for table orderinfo.
type Orderinfo struct {
	Orderid     string      `json:"orderid"     ` // 订单id
	Serverid    int         `json:"serverid"    ` // 服务ID
	Platn       string      `json:"platn"       ` // 角色渠道
	Roleid      int64       `json:"roleid"      ` // 角色ID
	Goodsid     int         `json:"goodsid"     ` // 商品ID
	PaymoneyCNY int         `json:"paymoneyCNY" ` // 充值金额_人民币
	PaymoneyUSD int         `json:"paymoneyUSD" ` // 充值金额_美元
	Currency    string      `json:"currency"    ` // 币种
	Ctime       *gtime.Time `json:"ctime"       ` // 订单时间
	Day         int         `json:"day"         ` // 订单时间day
	Hour        int         `json:"hour"        ` // 订单时间hour
	Month       int         `json:"month"       ` // 订单时间month
}
