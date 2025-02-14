// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this files as you wish.
// =================================================================================

package dao

import (
	"suask/internal/dao/internal"
)

// internalAttachmentsDao is internal type for wrapping internal DAO implements.
type internalAttachmentsDao = *internal.AttachmentsDao

// attachmentsDao is the data access object for table attachments.
// You can define custom methods on it to extend its functionality as you wish.
type attachmentsDao struct {
	internalAttachmentsDao
}

var (
	// Attachments is globally public accessible object for table attachments operations.
	Attachments = attachmentsDao{
		internal.NewAttachmentsDao(),
	}
)

// Fill with you ideas below.
