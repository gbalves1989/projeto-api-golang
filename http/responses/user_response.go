package responses

type UserData struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

type UserResponse struct {
	Code       int      `json:"code"`
	StatusCode string   `json:"status_code"`
	Message    string   `json:"message"`
	Data       UserData `json:"data,omitempty"`
}

type UserListResponse struct {
	Code       int                `json:"code"`
	StatusCode string             `json:"status_code"`
	Message    string             `json:"message"`
	Data       PaginationResponse `json:"data,omitempty"`
}
