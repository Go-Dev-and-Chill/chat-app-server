package models

// The `ServerProfile` type defines the structure of server profiles with fields for server ID, user
// ID, and user server role.
// @property {int} ServerID - ServerID is an integer that represents the unique identifier of a server.
// @property {int} UserID - UserID is an integer that represents the unique identifier of a user in the
// server profile.
// @property {string} UserServerRole - The `UserServerRole` property in the `ServerProfile` struct
// represents the role of a user within a specific server. This could be a designation such as "admin",
// "moderator", "member", or any other role that defines the user's permissions and responsibilities
// within that server.
type ServerProfile struct {
	ServerID       int    `json:"server_id"`
	UserID         int    `json:"user_id"`
	UserServerRole string `json:"user_server_role"`
}
