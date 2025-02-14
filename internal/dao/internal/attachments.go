// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AttachmentsDao is the data access object for table attachments.
type AttachmentsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns AttachmentsColumns // columns contains all the column names of Table for convenient usage.
}

// AttachmentsColumns defines and stores column names for table attachments.
type AttachmentsColumns struct {
	Id         string // 附件ID
	QuestionId string // 问题ID
	AnswerId   string // 回答ID
	Type       string // 附件类型（目前仅支持图片）
	FileId     string // 文件ID
}

// attachmentsColumns holds the columns for table attachments.
var attachmentsColumns = AttachmentsColumns{
	Id:         "id",
	QuestionId: "question_id",
	AnswerId:   "answer_id",
	Type:       "type",
	FileId:     "file_id",
}

// NewAttachmentsDao creates and returns a new DAO object for table data access.
func NewAttachmentsDao() *AttachmentsDao {
	return &AttachmentsDao{
		group:   "default",
		table:   "attachments",
		columns: attachmentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AttachmentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AttachmentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AttachmentsDao) Columns() AttachmentsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AttachmentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AttachmentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AttachmentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
