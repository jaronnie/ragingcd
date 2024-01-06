package query

type PageQuery struct {
	PageSize int `form:"pageSize"`
	PageNum  int `form:"pageNum"`
}
