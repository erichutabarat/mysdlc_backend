package model

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data	interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}