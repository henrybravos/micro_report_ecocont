package repositories

type Pagination struct {
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	TotalCount int    `json:"totalCount"`
	TotalPages int    `json:"totalPages"`
	Err        string `json:"err"`
}

type PaginationParams struct {
	Pagination bool
	Offset     int
	Limit      int
}
