package responses

type TokenResponse struct {
	Code        int    `json:"code"`
	StatusCode  string `json:"status_code"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}
