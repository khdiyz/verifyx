package models

type ListResponse struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"message"`
}

type Pagination struct {
	Page       int `json:"page"  default:"1"`
	Limit      int `json:"limit" default:"10"`
	Offset     int `json:"-" default:"0"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
}
