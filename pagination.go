package schemas

// PaginatedQueryResponse is a common schema that adheres and can be used with any form that binds using the directive for `form` fields
// all fields have `json` set as pascal case.
type PaginatedQueryResponse struct {
	Page     int    `json:"page" form:"page"`         // The page number to retrieve starting with 0 as the first page. Default is 0.
	PageSize int    `json:"pageSize" form:"pageSize"` // The number of items to retrieve per page. The max recommended page size is 100. Default is 10.
	Pages    int    `json:"pages" form:"pages"`       // The total number of pages available given the query. Default is 0.
	Total    int    `json:"total" form:"total"`       // The total number of items available given the query. Default is 0.
	Next     string `json:"next" form:"next"`         // The URL to retrieve the next page of results. Default is an empty string.
	Previous string `json:"previous" form:"previous"` // The URL to retrieve the previous page of results. Default is an empty string.
	Entities interface{}
}

// PaginatedQuery is a common schema that can be bound either for a query or a body of a HTTP request. All fields are bound to `json` as pascal case or
// can be bound from a form with the `form` dectorator
type PaginatedQuery struct {
	Page     int      `json:"page" form:"page"`         // Page of results to retrieve. Default is 0.
	PageSize int      `json:"pageSize" form:"pageSize"` // Number of results to retrieve per page. Default is 10. The max recommended is 100.
	Include  []string `json:"include" form:"include"`   // List of fields to include in the response. Default is all fields.
}
