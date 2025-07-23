package requests

type PaginationRequest struct {
	Limit  int `form:"limit,default=10" json:"limit"`
	Offset int `form:"offset,default=0" json:"offset"`
}
