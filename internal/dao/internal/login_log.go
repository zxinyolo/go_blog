// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LoginLogDao is the data access object for table login_log.
type LoginLogDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns LoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// LoginLogColumns defines and stores column names for table login_log.
type LoginLogColumns struct {
	Id          string //
	Username    string // 用户名称
	Ip          string // ip
	IpSource    string // ip来源
	Os          string // 操作系统
	Browser     string // 浏览器
	Status      string // 登录状态
	Description string // 操作描述
	CreateTime  string // 登录时间
	UserAgent   string // user-agent用户代理
}

//  loginLogColumns holds the columns for table login_log.
var loginLogColumns = LoginLogColumns{
	Id:          "id",
	Username:    "username",
	Ip:          "ip",
	IpSource:    "ip_source",
	Os:          "os",
	Browser:     "browser",
	Status:      "status",
	Description: "description",
	CreateTime:  "create_time",
	UserAgent:   "user_agent",
}

// NewLoginLogDao creates and returns a new DAO object for table data access.
func NewLoginLogDao() *LoginLogDao {
	return &LoginLogDao{
		group:   "default",
		table:   "login_log",
		columns: loginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LoginLogDao) Columns() LoginLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}