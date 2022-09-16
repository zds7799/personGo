// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-09-16 10:18:26
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PrimaryCachedatamonthDao is the data access object for table cachedatamonth.
type PrimaryCachedatamonthDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns PrimaryCachedatamonthColumns // columns contains all the column names of Table for convenient usage.
}

// PrimaryCachedatamonthColumns defines and stores column names for table cachedatamonth.
type PrimaryCachedatamonthColumns struct {
	Serverid    string // 服务ID
	Platn       string // 角色渠道
	Roletotal   string // 月总用户
	TotalDau    string // 月(每日)总活跃
	Month       string // month
	Mau         string // 月活跃MAU
	Newrole     string // 新用户
	Payroles    string // 当日付费用户数
	PaymoneyCny string // 当日付费总金额_人民币
	PaymoneyEny string // 当日付费总金额_美元
}

// primaryCachedatamonthColumns holds the columns for table cachedatamonth.
var primaryCachedatamonthColumns = PrimaryCachedatamonthColumns{
	Serverid:    "serverid",
	Platn:       "platn",
	Roletotal:   "roletotal",
	TotalDau:    "total_dau",
	Month:       "month",
	Mau:         "mau",
	Newrole:     "newrole",
	Payroles:    "payroles",
	PaymoneyCny: "paymoney_cny",
	PaymoneyEny: "paymoney_eny",
}

// NewPrimaryCachedatamonthDao creates and returns a new DAO object for table data access.
func NewPrimaryCachedatamonthDao() *PrimaryCachedatamonthDao {
	return &PrimaryCachedatamonthDao{
		group:   "default",
		table:   "cachedatamonth",
		columns: primaryCachedatamonthColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PrimaryCachedatamonthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PrimaryCachedatamonthDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PrimaryCachedatamonthDao) Columns() PrimaryCachedatamonthColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PrimaryCachedatamonthDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PrimaryCachedatamonthDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PrimaryCachedatamonthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
