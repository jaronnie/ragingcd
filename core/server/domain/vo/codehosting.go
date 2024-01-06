package vo

type CodeHostingVo struct {
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

type CodeHostingTableData struct {
	Total int              `json:"total"`
	Rows  []*CodeHostingVo `json:"rows"`
}
