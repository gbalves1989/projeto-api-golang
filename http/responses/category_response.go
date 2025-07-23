package responses

type CategoryData struct {
	ID          uint                     `json:"id"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Products    []ProductsByCategoryData `json:"products"`
}

type CategoryListData struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductsByCategoryData struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
}

type CategoryResponse struct {
	Code       int         `json:"code"`
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type CategoryListResponse struct {
	Code       int                `json:"code"`
	StatusCode string             `json:"status_code"`
	Message    string             `json:"message"`
	Data       PaginationResponse `json:"data"`
}
