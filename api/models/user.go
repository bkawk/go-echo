package models

type User struct {
	ID                 string `json:"id" bson:"_id" validate:"required"`
	Email              string `json:"email" bson:"email" validate:"required,email,max=100"`
	Username           string `json:"username" bson:"username" validate:"min=4,max=12"`
	Name	 	       string `json:"name" bson:"name" validate:"min=1,max=20"`
	Password           string `json:"password" bson:"password" validate:"max=64,min=8"`
	RefreshToken       string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	CreatedAt          int64  `json:"createdAt" bson:"createdAt"`
	VerificationCode   string `json:"verificationCode,omitempty" bson:"verificationCode,omitempty"`     // Verification code for email verification
	PasswordResetToken string `json:"passwordResetToken,omitempty" bson:"passwordResetToken,omitempty"` // Verification code for password reset
	LastSeen           int64  `json:"lastSeen,omitempty" bson:"lastSeen,omitempty"`
	IsVerified         bool   `bson:"isVerified"`
	ForgotPassword     int64  `json:"forgotPassword,omitempty" bson:"forgotPassword,omitempty"` // Timestamp for forgot password
}
