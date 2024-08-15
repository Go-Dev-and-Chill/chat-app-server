package models

type Profile struct {
	ProfileId     int    `json:"Profile_id"`
	UserID        int    `json:"user_id"`
	UserFirstName string `json:"User_first_name"`
	UserLastName  string `json:"User_last_name"`
	UserEmail     string `json:"User_email"`
}
