package requests

type ProductCreateOrUpdateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Stock       int     `json:"stock" binding:"required,min=0"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type ProductUploadRequest struct {
	ImageURL string `json:"image_url"`
}
