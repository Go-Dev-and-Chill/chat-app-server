package models

// The `ChannelProfile` type represents the profile of a user within a channel, including their ID,
// channel ID, and role.
// @property {int} ChannelID - The `ChannelID` property in the `ChannelProfile` struct represents the
// unique identifier for a channel.
// @property {int} UserID - The `UserID` property in the `ChannelProfile` struct represents the unique
// identifier of the user associated with the channel profile. It is an integer value used to identify
// the user within the system.
// @property {string} UserChannelRole - The `UserChannelRole` property in the `ChannelProfile` struct
// represents the role of a user within a specific channel. This role could indicate whether the user
// is an admin, a moderator, a member, or any other designated role within the channel.
type ChannelProfile struct {
	ChannelID       int    `json:"channel_id"`
	UserID          int    `json:"user_id"`
	UserChannelRole string `json:"user_channel_role"`
}
