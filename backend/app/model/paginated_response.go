package model

type PaginatedResponse[T any] struct {
	Data     []T `json:"data" example:"[{\"id\":220000,\"name\":\"Le Hong Phuc\",\"email\":\"hongphucle1010@gmail.com\",\"dob\":\"2004-10-10T00:00:00Z\",\"gpa\":4.0}]"`
	Total    int `json:"total" example:"1"`
	Page     int `json:"page" example:"1"`
	PageSize int `json:"page_size" example:"10"`
}
