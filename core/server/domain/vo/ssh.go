package vo

type SshVo struct {
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	Name     string
	Type     string // 登录方式 Password or SecretKey
	IP       string `json:"ip"`
	Port     int
	Username string
}

type SshTableData struct {
	Total int              `json:"total"`
	Rows  []*CodeHostingVo `json:"rows"`
}
