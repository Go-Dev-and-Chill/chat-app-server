package models

// The `Profile` struct defines a data structure with fields for profile information such as ID, user
// ID, first name, last name, and email.
// @property {int} ProfileId - The `ProfileId` field in the `Profile` struct represents the unique
// identifier for a profile.
// @property {int} UserID - UserID is an integer that represents the unique identifier of a user in the
// profile.
// @property {string} UserFirstName - User's first name
// @property {string} UserLastName - UserLastName is a field in the Profile struct that represents the
// last name of a user in a profile.
// @property {string} UserEmail - The `Profile` struct you provided represents a user profile with the
// following properties:
type Profile struct {
	ProfileId     int    `json:"Profile_id"`
	UserID        int    `json:"user_id"`
	UserFirstName string `json:"User_first_name"`
	UserLastName  string `json:"User_last_name"`
	UserEmail     string `json:"User_email"`
}
