package po

type CodeHosting struct {
	ID       int    `xorm:"pk autoincr COMMENT('id') 'id'"`
	UUID     string `xorm:"COMMENT('uuid') 'uuid'"`
	Name     string `xorm:"COMMENT(' 名称') 'name'"`
	Type     string `xorm:"COMMENT('类型 github gitlab gitea local') 'type'"`
	Url      string `xorm:"COMMENT('url') 'url'"`
	Username string `xorm:"COMMENT('用户名') 'username'"`
	Token    string `xorm:"COMMENT('token') 'token'"`

	// auth
	UserID int `xorm:"COMMENT('鉴权-user_id') 'user_id'"`
}
