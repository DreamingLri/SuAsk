// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachersDao is the data access object for table teachers.
type TeachersDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TeachersColumns // columns contains all the column names of Table for convenient usage.
}

// TeachersColumns defines and stores column names for table teachers.
type TeachersColumns struct {
	Id           string //
	Responses    string // 回复数
	Name         string // 老师名字
	AvatarUrl    string // 老师头像链接
	Introduction string // 老师简介
	Email        string // 老师邮箱
	Perm         string // 提问箱权限
}

// teachersColumns holds the columns for table teachers.
var teachersColumns = TeachersColumns{
	Id:           "id",
	Responses:    "responses",
	Name:         "name",
	AvatarUrl:    "avatar_url",
	Introduction: "introduction",
	Email:        "email",
	Perm:         "perm",
}

// NewTeachersDao creates and returns a new DAO object for table data access.
func NewTeachersDao() *TeachersDao {
	return &TeachersDao{
		group:   "default",
		table:   "teachers",
		columns: teachersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TeachersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TeachersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TeachersDao) Columns() TeachersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TeachersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TeachersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TeachersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
