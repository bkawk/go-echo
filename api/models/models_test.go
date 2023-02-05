package models

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestUserStruct(t *testing.T) {
	validate := validator.New()

	testCases := []struct {
		user  User
		isErr bool
	}{
		{User{ID: "", Email: "test@test.com", Username: "testuser", Password: "testpassword", CreatedAt: 123456789, LastSeen: 123456789}, true},
		{User{ID: "testid", Email: "", Username: "testuser", Password: "testpassword", CreatedAt: 123456789, LastSeen: 123456789}, true},
		{User{ID: "testid", Email: "testtest.com", Username: "testuser", Password: "testpassword", CreatedAt: 123456789, LastSeen: 123456789}, true},
		{User{ID: "testid", Email: "test@test.com", Username: "", Password: "testpassword", CreatedAt: 123456789, LastSeen: 123456789}, true},
		{User{ID: "testid", Email: "test@test.com", Username: "test", Password: "", CreatedAt: 123456789, LastSeen: 123456789}, true},
		{User{ID: "testid", Email: "test@test.com", Username: "testuser", Password: "testpassword", CreatedAt: 0, LastSeen: 123456789}, false},
		{User{ID: "testid", Email: "test@test.com", Username: "testuser", Password: "testpassword", CreatedAt: 123456789, LastSeen: 0}, false},
	}

	for _, tc := range testCases {
		err := validate.Struct(tc.user)
		if (err == nil) == tc.isErr {
			t.Errorf("Expected error: %v, but got: %v", tc.isErr, err)
		}
	}
}
