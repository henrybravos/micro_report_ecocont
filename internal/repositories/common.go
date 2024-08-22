package repositories

type Pagination struct {
	Page       int
	PageSize   int
	TotalCount int
	TotalPages int
}
type PaginationResponse struct {
	Data       []interface{}
	Pagination Pagination
}
type PaginationParams struct {
	Pagination bool
	Offset     int
	Limit      int
}
