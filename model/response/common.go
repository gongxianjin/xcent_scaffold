package response

type PageResult struct {
	List     interface{} `json:"list"`
	TotalCount    int64       `json:"totalCount"`
	PageNo     int         `json:"pageNo"`
	PageSize int         `json:"pageSize"`
}
