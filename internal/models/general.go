package models

type ListResponse struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"message"`
}

type Pagination struct {
	Page       int64 `json:"page"  default:"1"`
	Limit      int64 `json:"limit" default:"10"`
	Offset     int64 `json:"-" default:"0"`
	PageCount  int64 `json:"page_count"`
	TotalCount int64 `json:"total_count"`
}
