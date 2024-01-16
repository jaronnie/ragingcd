package vo

type SshVo struct {
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Type     string `json:"type"` // 登录方式 password or private_key
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}

type SshTableData struct {
	Total int              `json:"total"`
	Rows  []*CodeHostingVo `json:"rows"`
}
