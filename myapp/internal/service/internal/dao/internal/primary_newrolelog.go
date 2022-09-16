// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-09-16 10:18:26
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PrimaryNewrolelogDao is the data access object for table newrolelog.
type PrimaryNewrolelogDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns PrimaryNewrolelogColumns // columns contains all the column names of Table for convenient usage.
}

// PrimaryNewrolelogColumns defines and stores column names for table newrolelog.
type PrimaryNewrolelogColumns struct {
	Roleid   string // 角色id
	Serverid string // 服务ID
	Platn    string // 角色渠道
	Ctime    string // 创建时间
	Day      string // 创建时间day
	Hour     string // 创建时间hour
	Month    string // 创建时间month
}

// primaryNewrolelogColumns holds the columns for table newrolelog.
var primaryNewrolelogColumns = PrimaryNewrolelogColumns{
	Roleid:   "roleid",
	Serverid: "serverid",
	Platn:    "platn",
	Ctime:    "ctime",
	Day:      "day",
	Hour:     "hour",
	Month:    "month",
}

// NewPrimaryNewrolelogDao creates and returns a new DAO object for table data access.
func NewPrimaryNewrolelogDao() *PrimaryNewrolelogDao {
	return &PrimaryNewrolelogDao{
		group:   "default",
		table:   "newrolelog",
		columns: primaryNewrolelogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PrimaryNewrolelogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PrimaryNewrolelogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PrimaryNewrolelogDao) Columns() PrimaryNewrolelogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PrimaryNewrolelogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PrimaryNewrolelogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PrimaryNewrolelogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}