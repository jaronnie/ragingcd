package po

func init() {
	register(Docker{})
}

type Docker struct {
	Base `xorm:"extends"`

	ContainerName string `xorm:"COMMENT('容器名称') 'container_name'"`
	ContainerID   string `xorm:"COMMENT('id') 'container_id'"`
}
