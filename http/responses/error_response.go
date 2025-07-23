package responses

type ErrorResponse struct {
	Code       int    `json:"code"`
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}
