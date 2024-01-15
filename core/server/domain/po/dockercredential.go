package po

func init() {
	register(DockerCredential{})
}

type DockerCredential struct {
	Base `xorm:"extends"`

	Name string `xorm:"COMMENT('名称') 'name'"`
	IP   string `xorm:"COMMENT('ip') 'ip'"`
	Port int    `xorm:"COMMENT('端口') 'port'"`
	Type string `xorm:"COMMENT('类型 ssh tcp') 'type'"`
}
