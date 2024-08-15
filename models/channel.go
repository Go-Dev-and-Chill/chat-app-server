package models

type Channel struct {
	ChannelID   int    `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	ChannelDesc string `json:"channel_desc"`
}
