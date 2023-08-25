package common

import "time"

//model structs used for both input and output

type QueryParams struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Query    string `json:"query"`
	Filter   string `json:"filter"`
	SortBy   string `json:"sort_by"`
	SortDesc bool   `json:"sort_by"`
}

type SalesReportDateRange struct {
	StartDate  time.Time  `json:"start_date"`
	EndDate    time.Time  `json:"end_date"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Count      int `json:"count"`
	PageNumber int `json:"page_number"`
}
