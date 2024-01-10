package po

import "time"

type Base struct {
	CreatedAt time.Time `xorm:"created comment('创建时间，自动添加') index"`
	UpdatedAt time.Time `xorm:"updated comment('更新时间，自动更新') index"`
}
