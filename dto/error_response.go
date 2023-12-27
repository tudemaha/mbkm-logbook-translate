package dto

type ErrorResponse struct {
	Error Error       `json:"error"`
	Meta  interface{} `json:"meta"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
