package web

type Pagination struct {
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
	RowCount  int `json:"row_count"`
	PageCount int `json:"page_count"`
}
