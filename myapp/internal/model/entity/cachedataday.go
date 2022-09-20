// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-09-19 15:24:37
// =================================================================================

package entity

// Cachedataday is the golang structure for table cachedataday.
type Cachedataday struct {
	Serverid       int    `json:"serverid"       ` // 服务ID
	Platn          string `json:"platn"          ` // 角色渠道
	Day            int    `json:"day"            ` // day
	Dau            int    `json:"dau"            ` // 日活跃
	Newrolecount   int    `json:"newrolecount"   ` // 新用户
	Payroles       int    `json:"payroles"       ` // 当日付费用户数
	PaymoneyCny    int    `json:"paymoneyCny"    ` // 当日付费总金额_人民币
	PaymoneyEny    int    `json:"paymoneyEny"    ` // 当日付费总金额_美元
	Newpayroles    int    `json:"newpayroles"    ` // 当日新用户付费数
	NewpaymoneyCny int    `json:"newpaymoneyCny" ` // 当日新用户付费总金额_人民币
	NewpaymoneyEny int    `json:"newpaymoneyEny" ` // 当日新用户付费总金额_美元
}
