package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Authentication : "using for auth"
type Authentication struct {
	UserID   primitive.ObjectID `json:"userID,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName string             `json:"username,omitempty"`
	Status   string             `json:"status,omitempty"`
	Role     string             `json:"role,omitempty"`
}

//Token : "struct"
type Token struct {
	Token string `json:"token"`
}

//Login : "login"
type Login struct {
	UserType string   `json:"userType"`
	UserName string   `json:"userName"`
	PassWord string   `json:"password"`
	Location Location `json:"location,omitempty" bson:"location,omitempty"`
}

//OTPLogin : ""
type OTPLogin struct {
	UserType string `json:"userType"`
	Mobile   string `json:"mobile"`
	OTP      string `json:"otp"`
	Scenario string `json:"scenario"`
}
