// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-09-19 15:24:37
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CachedatadayDao is the data access object for table cachedataday.
type CachedatadayDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns CachedatadayColumns // columns contains all the column names of Table for convenient usage.
}

// CachedatadayColumns defines and stores column names for table cachedataday.
type CachedatadayColumns struct {
	Serverid       string // 服务ID
	Platn          string // 角色渠道
	Day            string // day
	Dau            string // 日活跃
	Newrolecount   string // 新用户
	Payroles       string // 当日付费用户数
	PaymoneyCny    string // 当日付费总金额_人民币
	PaymoneyEny    string // 当日付费总金额_美元
	Newpayroles    string // 当日新用户付费数
	NewpaymoneyCny string // 当日新用户付费总金额_人民币
	NewpaymoneyEny string // 当日新用户付费总金额_美元
}

// cachedatadayColumns holds the columns for table cachedataday.
var cachedatadayColumns = CachedatadayColumns{
	Serverid:       "serverid",
	Platn:          "platn",
	Day:            "day",
	Dau:            "dau",
	Newrolecount:   "newrolecount",
	Payroles:       "payroles",
	PaymoneyCny:    "paymoney_cny",
	PaymoneyEny:    "paymoney_eny",
	Newpayroles:    "newpayroles",
	NewpaymoneyCny: "newpaymoney_cny",
	NewpaymoneyEny: "newpaymoney_eny",
}

// NewCachedatadayDao creates and returns a new DAO object for table data access.
func NewCachedatadayDao() *CachedatadayDao {
	return &CachedatadayDao{
		group:   "default",
		table:   "cachedataday",
		columns: cachedatadayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CachedatadayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CachedatadayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CachedatadayDao) Columns() CachedatadayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CachedatadayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CachedatadayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CachedatadayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}