package po

func init() {
	register(Ssh{})
}

type Ssh struct {
	Base `xorm:"extends"`

	Name       string `xorm:"COMMENT('名称') 'name'"`
	IP         string `xorm:"COMMENT('ip') 'ip'"`
	Port       int    `xorm:"COMMENT('端口') 'port'"`
	Type       string `xorm:"COMMENT('类型 password private_key') 'type'"`
	Username   string `xorm:"COMMENT('用户名') 'username'"`
	Password   string `xorm:"COMMENT('密码') 'password'"`
	PrivateKey string `xorm:"COMMENT('私钥') 'private_key'"`
}
