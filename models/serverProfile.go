package models

type ServerProfile struct {
	ServerID       int    `json:"server_id"`
	UserID         int    `json:"user_id"`
	UserServerRole string `json:"user_server_role"`
}
