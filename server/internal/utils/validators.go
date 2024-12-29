package utils

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	uppercaseRegex   = regexp.MustCompile(`[A-Z]`)
	lowercaseRegex   = regexp.MustCompile(`[a-z]`)
	digitRegex       = regexp.MustCompile(`\d`)
	specialCharRegex = regexp.MustCompile(`[#?!@$%^&*-]`)
	minLength        = 8
	maxLength        = 32
)

// IsValidPassword reports whether the password is valid.
//
// A valid password must contain:
//   - Minimum eight characters,
//   - At most thirty-two characters,
//   - At least one uppercase letter,
//   - At least one lowercase letter,
//   - At least one number,
//   - At least one special character.
func IsValidPassword(password string) bool {
	if len(password) < minLength || len(password) > maxLength {
		return false
	}
	if !uppercaseRegex.MatchString(password) {
		return false
	}
	if !lowercaseRegex.MatchString(password) {
		return false
	}
	if !digitRegex.MatchString(password) {
		return false
	}
	if !specialCharRegex.MatchString(password) {
		return false
	}
	return true
}

var validImageExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

// HasImageValidExtension reports whether the image URL extension is valid.
//
// A valid image extension must be one of the following:
//   - .jpg,
//   - .jpeg,
//   - .png,
//   - .gif,
//   - .webp.
func HasImageValidExtension(imageURL string) bool {
	imageURL = strings.ToLower(imageURL)
	for extension := range validImageExtensions {
		if strings.HasSuffix(imageURL, extension) {
			return true
		}
	}
	return false
}

func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func ContainsEmptySpace(s string) bool {
	return strings.Contains(s, " ")
}

func IsMalformedURL(s string) bool {
	parsedURL, err := url.Parse(s)
	return err == nil && parsedURL.Host == ""
}

func IsInvalidProtocol(s string) bool {
	parsedURL, err := url.Parse(s)
	return err != nil || parsedURL.Scheme != "http" && parsedURL.Scheme != "https"
}

func IsImageLiked(images []string, imageURL string) bool {
	for _, img := range images {
		if img == imageURL {
			return true
		}
	}
	return false
}
