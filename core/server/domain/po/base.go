package po

import "time"

type Base struct {
	CreatedAt time.Time `xorm:"created comment('创建时间，自动添加') index"`
	UpdatedAt time.Time `xorm:"updated comment('更新时间，自动更新') index"`
	ID        int       `xorm:"pk autoincr COMMENT('id') 'id'"`
	UUID      string    `xorm:"COMMENT('uuid') 'uuid'"`

	// auth
	UserID int `xorm:"COMMENT('鉴权-user_id') 'user_id'"`
}
