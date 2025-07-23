package requests

type CategoryCreateOrUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
}
