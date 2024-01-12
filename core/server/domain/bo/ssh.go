package bo

type CreateSshBo struct {
	Name       string
	IP         string `json:"ip"`
	Port       int
	Type       string // 登录方式 password or private_key
	Username   string
	Password   string // 密码
	PrivateKey string `json:"private_key"` // 私钥
}
