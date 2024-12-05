// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"suask/internal/dao/internal"
)

// internalQuestionsDao is internal type for wrapping internal DAO implements.
type internalQuestionsDao = *internal.QuestionsDao

// questionsDao is the data access object for table questions.
// You can define custom methods on it to extend its functionality as you wish.
type questionsDao struct {
	internalQuestionsDao
}

var (
	// Questions is globally public accessible object for table questions operations.
	Questions = questionsDao{
		internal.NewQuestionsDao(),
	}
)

// Fill with you ideas below.