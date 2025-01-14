// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Favorites is the golang structure for table favorites.
type Favorites struct {
	Id         int64       `json:"id"         orm:"id"          description:"收藏（置顶）ID"` // 收藏（置顶）ID
	UserId     int         `json:"userId"     orm:"user_id"     description:"用户ID"`     // 用户ID
	QuestionId int         `json:"questionId" orm:"question_id" description:"问题ID"`     // 问题ID
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`     // 创建时间
	Package    string      `json:"package"    orm:"package"     description:"收藏夹"`      // 收藏夹
}
