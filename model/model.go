package model

type Status int

const (
	Success Status = iota
	Request
	Resend
	Invalid
	NotFound
	InternalError
	Unauthorized
	Forbidden
	Close
)

type Message struct {
	Content     string `json:"content"`
	JsonContent string `json:"json_content"`
	Error       string `json:"error"`
	Token       string `json:"token"`
	Status      Status `json:"status"`
}
