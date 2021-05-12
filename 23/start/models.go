package start

type HttpError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
