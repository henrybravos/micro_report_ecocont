package services

import v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"

func GetPaginationOrDefault(page, pageSize int32) *v1.Pagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 15
	} else if pageSize > 1000 {
		pageSize = 1000
	}
	offset := (page - 1) * pageSize
	return &v1.Pagination{
		Page:       page,
		PageSize:   pageSize,
		Offset:     offset,
		TotalPages: 0,
		TotalCount: 0,
	}
}
