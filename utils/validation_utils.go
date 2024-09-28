package utils

import (
	"regexp"
)

func IsValidUUID(uuid string) bool {
	// Regular expression for UUID validation
	r := regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")
	return r.MatchString(uuid)
}

// ... other validation functions (e.g., for email, postal code, etc.) ...
