package requests

type UserUpdateRequest struct {
	Name string `json:"name,omitempty"`
}

type UserUpdatePasswordRequest struct {
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6"`
}

type UserUploadRequest struct {
	AvatarURL string `json:"avatar_url,omitempty"`
}

type UserDeleteRequest struct {
	Email string `json:"email" binding:"required,email"`
}
