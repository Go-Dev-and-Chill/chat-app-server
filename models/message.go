package models

// The Message type defines a structure with fields for message ID, channel ID, and user ID.
// @property {int} MessageID - The `MessageID` field in the `Message` struct represents the unique
// identifier for a message.
// @property {int} ChannelID - The `ChannelID` property in the `Message` struct represents the unique
// identifier of the channel where the message is posted.
// @property {int} UserID - The `UserID` property in the `Message` struct represents the unique
// identifier of the user who sent the message. This identifier can be used to associate the message
// with a specific user in the system.
type Message struct {
	MessageID int `json:"message_id"`
	ChannelID int `json:"channel_id"`
	UserID    int `json:"user_id"`
}
