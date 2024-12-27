package utils_test

import (
	"server/internal/utils"
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{
			name:     "Valid password",
			password: "Password1!",
			want:     true,
		},
		{
			name:     "Invalid password - missing uppercase letter",
			password: "password1!",
			want:     false,
		},
		{
			name:     "Invalid password - missing lowercase letter",
			password: "PASSWORD1!",
			want:     false,
		},
		{
			name:     "Invalid password - missing number",
			password: "Password!",
			want:     false,
		},
		{
			name:     "Invalid password - missing special character",
			password: "Password1",
			want:     false,
		},
		{
			name:     "Invalid password - too short",
			password: "Pass1!",
			want:     false,
		},
		{
			name:     "Invalid password - too long",
			password: "Password1234567890123456789012345678901!",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsValidPassword(tt.password); got != tt.want {
				t.Errorf("IsValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
