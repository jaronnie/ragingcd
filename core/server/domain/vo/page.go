package vo

type PageData struct {
	Total int         `json:"total"`
	Rows  interface{} `json:"rows"`
}
