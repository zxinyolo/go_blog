// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VisitorDao is the data access object for table visitor.
type VisitorDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns VisitorColumns // columns contains all the column names of Table for convenient usage.
}

// VisitorColumns defines and stores column names for table visitor.
type VisitorColumns struct {
	Id         string //
	Uuid       string // 访客标识码
	Ip         string // ip
	IpSource   string // ip来源
	Os         string // 操作系统
	Browser    string // 浏览器
	CreateTime string // 首次访问时间
	LastTime   string // 最后访问时间
	Pv         string // 访问页数统计
	UserAgent  string // user-agent用户代理
}

//  visitorColumns holds the columns for table visitor.
var visitorColumns = VisitorColumns{
	Id:         "id",
	Uuid:       "uuid",
	Ip:         "ip",
	IpSource:   "ip_source",
	Os:         "os",
	Browser:    "browser",
	CreateTime: "create_time",
	LastTime:   "last_time",
	Pv:         "pv",
	UserAgent:  "user_agent",
}

// NewVisitorDao creates and returns a new DAO object for table data access.
func NewVisitorDao() *VisitorDao {
	return &VisitorDao{
		group:   "default",
		table:   "visitor",
		columns: visitorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VisitorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VisitorDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VisitorDao) Columns() VisitorColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VisitorDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VisitorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VisitorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}