package models

// The User type defines a structure with fields for UserID, Username, and Password.
// @property {int} UserID - The `UserID` field in the `User` struct represents the unique identifier
// for a user. It is of type `int` and is tagged with `json:"id"` for JSON marshaling purposes.
// @property {string} Username - The `Username` property in the `User` struct represents the username
// of a user.
// @property {string} Password - The `User` struct defines a user object with three properties:
// `UserID`, `Username`, and `Password`. The `Password` property is a string type field that stores the
// user's password.
type User struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
