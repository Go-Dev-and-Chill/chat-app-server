package models

// The `Channel` type represents a channel with an ID, name, and description.
// @property {int} ChannelID - The `ChannelID` field in the `Channel` struct represents the unique
// identifier for a channel. It is of type `int` and is tagged with `json:"channel_id"` for JSON
// marshaling and unmarshaling.
// @property {string} ChannelName - ChannelName is a string property that represents the name of a
// channel.
// @property {string} ChannelDesc - The `ChannelDesc` property in the `Channel` struct represents the
// description of a channel. It provides a brief explanation or summary of what the channel is about.
type Channel struct {
	ChannelID   int    `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	ChannelDesc string `json:"channel_desc"`
}
