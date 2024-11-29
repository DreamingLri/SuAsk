// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"suask/internal/dao/internal"
)

// internalFilesDao is internal type for wrapping internal DAO implements.
type internalFilesDao = *internal.FilesDao

// filesDao is the data access object for table files.
// You can define custom methods on it to extend its functionality as you wish.
type filesDao struct {
	internalFilesDao
}

var (
	// Files is globally public accessible object for table files operations.
	Files = filesDao{
		internal.NewFilesDao(),
	}
)

// Fill with you ideas below.
