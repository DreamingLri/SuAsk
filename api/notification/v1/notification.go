package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"suask/internal/model"
)

type NotificationGetReq struct {
	g.Meta `path:"/notification" method:"GET" tags:"Notification" summary:"通过用户Id拿到通知"`
	UserId int `json:"user_id" dc:"用户ID"`
}

type NotificationGetRes struct {
	Notifications []model.Notification `json:"notifications" dc:"通知列表"`
}

type NotificationUpdateReq struct {
	g.Meta `path:"/notification" method:"PUT" tags:"Notification" summary:"更新已读信息"`
	Id     int `json:"id" dc:"提醒ID"`
}

type NotificationUpdateRes struct {
	Id     int  `json:"id" dc:"提醒ID"`
	IsRead bool `json:"is_read" dc:"是否已读"`
}

type NotificationDeleteReq struct {
	g.Meta `path:"/notification" method:"DELETE" tags:"Notification" summary:"删除提醒"`
	Id     int `json:"id" dc:"提醒ID"`
}

type NotificationDeleteRes struct{}