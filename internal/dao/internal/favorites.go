// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FavoritesDao is the data access object for table favorites.
type FavoritesDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns FavoritesColumns // columns contains all the column names of Table for convenient usage.
}

// FavoritesColumns defines and stores column names for table favorites.
type FavoritesColumns struct {
	Id         string // 收藏（置顶）ID
	UserId     string // 用户ID
	QuestionId string // 问题ID
	CreatedAt  string // 创建时间
	Package    string // 收藏夹
}

// favoritesColumns holds the columns for table favorites.
var favoritesColumns = FavoritesColumns{
	Id:         "id",
	UserId:     "user_id",
	QuestionId: "question_id",
	CreatedAt:  "created_at",
	Package:    "package",
}

// NewFavoritesDao creates and returns a new DAO object for table data access.
func NewFavoritesDao() *FavoritesDao {
	return &FavoritesDao{
		group:   "default",
		table:   "favorites",
		columns: favoritesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FavoritesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FavoritesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FavoritesDao) Columns() FavoritesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FavoritesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FavoritesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FavoritesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
