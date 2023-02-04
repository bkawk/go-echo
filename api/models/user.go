package models

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID           string `json:"id" bson:"_id" validate:"required, startswith=usr_"`
	Email        string `json:"email" bson:"email" validate:"required,email, max=100"`
	Username     string `json:"username" bson:"username" validate:"required, min=4,max=12"`
	Password     string `json:"password" bson:"password" validate:"required, password, max=64 min=8"`
	RefreshToken string `json:"refreshToken" bson:"refreshToken"`
	CreatedAt    int64  `json:"createdAt" bson:"createdAt"`
	LastSeen     int64  `json:"lastSeen" bson:"lastSeen"`
}

func init() {
	validate := validator.New()
	validate.RegisterValidation("startswith", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^usr_`).MatchString(fl.Field().String())
	})
}
