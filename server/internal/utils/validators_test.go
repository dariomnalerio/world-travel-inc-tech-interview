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

func TestHasImageValidExtension(t *testing.T) {
	tests := []struct {
		name     string
		imageURL string
		want     bool
	}{
		{
			name:     "Valid extension - jpg",
			imageURL: "image.jpg",
			want:     true,
		},
		{
			name:     "Valid extension - jpeg",
			imageURL: "image.jpeg",
			want:     true,
		},
		{
			name:     "Valid extension - png",
			imageURL: "image.png",
			want:     true,
		},
		{
			name:     "Valid extension - gif",
			imageURL: "image.gif",
			want:     true,
		},
		{
			name:     "Valid extension - webp",
			imageURL: "image.webp",
			want:     true,
		},
		{
			name:     "Invalid extension - bmp",
			imageURL: "image.bmp",
			want:     false,
		},
		{
			name:     "Invalid extension - txt",
			imageURL: "image.txt",
			want:     false,
		},
		{
			name:     "Invalid extension - no extension",
			imageURL: "image",
			want:     false,
		},
		{
			name:     "Invalid extension - empty string",
			imageURL: "",
			want:     false,
		},
		{
			name:     "Invalid extension - no '.' in suffix",
			imageURL: "imagejpg",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.HasImageValidExtension(tt.imageURL); got != tt.want {
				t.Errorf("HasImageValidExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestIsEmptyString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Empty string",
			s:    "",
			want: true,
		},
		{
			name: "String with only spaces",
			s:    "   ",
			want: true,
		},
		{
			name: "Non-empty string",
			s:    "hello",
			want: false,
		},
		{
			name: "String with spaces and characters",
			s:    "  hello  ",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsEmptyString(tt.s); got != tt.want {
				t.Errorf("IsEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsEmptySpace(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "String with space",
			s:    "hello world",
			want: true,
		},
		{
			name: "String without space",
			s:    "helloworld",
			want: false,
		},
		{
			name: "Empty string",
			s:    "",
			want: false,
		},
		{
			name: "String with multiple spaces",
			s:    "hello  world",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ContainsEmptySpace(tt.s); got != tt.want {
				t.Errorf("ContainsEmptySpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMalformedURL(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Malformed URL - missing host",
			s:    "http://",
			want: true,
		},
		{
			name: "Malformed URL - missing scheme",
			s:    "www.example.com",
			want: true,
		},
		{
			name: "Valid URL",
			s:    "http://www.example.com",
			want: false,
		},
		{
			name: "Valid URL with path",
			s:    "http://www.example.com/path",
			want: false,
		},
		{
			name: "Empty string",
			s:    "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsMalformedURL(tt.s); got != tt.want {
				t.Errorf("IsMalformedURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInvalidProtocol(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Invalid protocol - ftp",
			s:    "ftp://www.example.com",
			want: true,
		},
		{
			name: "Invalid protocol - mailto",
			s:    "mailto:someone@example.com",
			want: true,
		},
		{
			name: "Valid protocol - http",
			s:    "http://www.example.com",
			want: false,
		},
		{
			name: "Valid protocol - https",
			s:    "https://www.example.com",
			want: false,
		},
		{
			name: "Empty string",
			s:    "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsInvalidProtocol(tt.s); got != tt.want {
				t.Errorf("IsInvalidProtocol() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestIsImageLiked(t *testing.T) {
	tests := []struct {
		name     string
		images   []string
		imageURL string
		want     bool
	}{
		{
			name:     "Image is liked",
			images:   []string{"image1.jpg", "image2.png", "image3.gif"},
			imageURL: "image2.png",
			want:     true,
		},
		{
			name:     "Image is not liked",
			images:   []string{"image1.jpg", "image2.png", "image3.gif"},
			imageURL: "image4.webp",
			want:     false,
		},
		{
			name:     "Empty images list",
			images:   []string{},
			imageURL: "image1.jpg",
			want:     false,
		},
		{
			name:     "Empty imageURL",
			images:   []string{"image1.jpg", "image2.png", "image3.gif"},
			imageURL: "",
			want:     false,
		},
		{
			name:     "Both images list and imageURL are empty",
			images:   []string{},
			imageURL: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsImageLiked(tt.images, tt.imageURL); got != tt.want {
				t.Errorf("IsImageLiked() = %v, want %v", got, tt.want)
			}
		})
	}
}
