package start

type HttpError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
