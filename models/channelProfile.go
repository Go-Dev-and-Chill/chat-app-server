package models

type ChannelProfile struct {
	ChannelID       int    `json:"channel_id"`
	UserID          int    `json:"user_id"`
	UserChannelRole string `json:"user_channel_role"`
}
