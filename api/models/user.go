package models

type User struct {
	ID               string `json:"id" bson:"_id" validate:"required"`
	Email            string `json:"email" bson:"email" validate:"required,email,max=100"`
	Username         string `json:"username" bson:"username" validate:"required,min=4,max=12"`
	Password         string `json:"password" bson:"password" validate:"required,max=64,min=8"`
	RefreshToken     string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	CreatedAt        int64  `json:"createdAt" bson:"createdAt"`
	VerificationCode int    `json:"verification,omitempty" bson:"verification,omitempty"`
	LastSeen         int64  `json:"lastSeen,omitempty" bson:"lastSeen,omitempty"`
	IsVerified       bool   `bson:"isVerified"`
}
