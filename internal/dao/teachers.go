// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"suask/internal/dao/internal"
)

// internalTeachersDao is internal type for wrapping internal DAO implements.
type internalTeachersDao = *internal.TeachersDao

// teachersDao is the data access object for table teachers.
// You can define custom methods on it to extend its functionality as you wish.
type teachersDao struct {
	internalTeachersDao
}

var (
	// Teachers is globally public accessible object for table teachers operations.
	Teachers = teachersDao{
		internal.NewTeachersDao(),
	}
)

// Fill with you ideas below.