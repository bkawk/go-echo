package models

type User struct {
	Email    string `json:"email" bson:"email" validate:"required,email, max=100"`
	Username string `json:"username" bson:"username" validate:"required, min=4,max=12"`
	Password string `json:"password" bson:"password" validate:"required, password, max=64 min=8"`
}
