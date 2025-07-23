package responses

type PaginationResponse struct {
	Total       int64       `json:"total"`
	Limit       int         `json:"limit"`
	Offset      int         `json:"offset"`
	Data        interface{} `json:"data"`
	HasNext     bool        `json:"hasNext"`
	HasPrevious bool        `json:"hasPrevious"`
}
