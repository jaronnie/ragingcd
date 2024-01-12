package po

func init() {
	register(CodeHosting{})
}

type CodeHosting struct {
	Base `xorm:"extends"`

	Name     string `xorm:"COMMENT(' 名称') 'name'"`
	Type     string `xorm:"COMMENT('类型 github gitlab gitea local') 'type'"`
	Url      string `xorm:"COMMENT('url') 'url'"`
	Username string `xorm:"COMMENT('用户名') 'username'"`
	Token    string `xorm:"COMMENT('token') 'token'"`
}
